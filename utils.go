package main

import (
	"encoding/json"
	"crypto/x509/pkix"
	"errors"
	"fmt"
	"time"
	"crypto/x509"
	"encoding/pem"
	"encoding/base64"
	"encoding/asn1"
	"encoding/hex"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/mitchellh/mapstructure"
	"github.com/xeipuuv/gojsonschema"
)

func ParseInput(inputJSON []byte, result interface{}, schema string) error {
	methodName := "[ParseInput]"
	var input interface{}

	err := validateJSONSchema(schema, string(inputJSON))
	if err != nil {
		return errors.New(methodName + " " + err.Error())
	}

	err = json.Unmarshal(inputJSON, &input)
	if err != nil {
		return err
	}

	//Decoding Map to Struct
	config := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeHookFunc(time.RFC3339),
		WeaklyTypedInput: true,
		Result:           result,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.Decode(input)
	if err != nil {
		return err
	}

	return nil
}

func validateJSONSchema(schema string, jsonObj string) error {
	methodName := "[validateJSONSchema]"
	loaderSchema := gojsonschema.NewStringLoader(schema)
	loaderObj := gojsonschema.NewStringLoader(jsonObj)
	result, err := gojsonschema.Validate(loaderSchema, loaderObj)
	//logger.new("schema : " + schema)
	if err != nil {
		logger.Errorf(methodName + err.Error())
		return errors.New(methodName + err.Error())
	}
	if !result.Valid() {

		errMsg := ""
		for _, err := range result.Errors() {
			errMsg = errMsg + " " + err.String()
		}
		logger.Errorf(methodName + " The document is not valid. see errors: " + errMsg)
		//return errors.New(methodName + " Can't be Validate JSON Schema. Error: " + errMsg + "schema : " + schema)
		return errors.New("jsonObj : " + jsonObj)
	}
	return nil
}

//retrieve enrollment id of user
func getEnrollmentID(stub shim.ChaincodeStubInterface) (string, error) {
	// // GetCreator returns marshaled serialized identity of the client
	serializedID, _ := stub.GetCreator()
	sID := &msp.SerializedIdentity{}
	err := proto.Unmarshal(serializedID, sID)
	if err != nil {
		return "", fmt.Errorf("Could not deserialize a SerializedIdentity, err %s", err)
	}
	bl, _ := pem.Decode(sID.IdBytes)
	if bl == nil {
		return "", fmt.Errorf("Failed to decode PEM structure")
	}
	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		return "", fmt.Errorf("Unable to parse certificate %s", err)
	}
	enrollmentID := string(cert.Subject.CommonName)

	logger.Debugf("enrollmentId: %s", enrollmentID)
	return enrollmentID, err
}

func getIDFromCert(cert string) (string, error) {
	methodName := "[utils.getIDFromCert()]"
	bl, _ := pem.Decode([]byte(cert))
	if bl == nil {
		return "", fmt.Errorf(methodName + " Failed to decode PEM structure. Received cert: " + cert)
	}

	parsedCert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
	 	return "", fmt.Errorf(methodName + " Unable to parse certificate " + err.Error())
	}

	idNotEncoded := fmt.Sprintf("x509::%s::%s", getDN(&parsedCert.Subject), getDN(&parsedCert.Issuer))
	certID := base64.StdEncoding.EncodeToString([]byte(idNotEncoded))

	return certID, nil
}

func getDN(name *pkix.Name) string {
	r := name.ToRDNSequence()
	s := ""
	var attributeTypeNames = map[string]string{
		"2.5.4.6":  "C",
		"2.5.4.10": "O",
		"2.5.4.11": "OU",
		"2.5.4.3":  "CN",
		"2.5.4.5":  "SERIALNUMBER",
		"2.5.4.7":  "L",
		"2.5.4.8":  "ST",
		"2.5.4.9":  "STREET",
		"2.5.4.17": "POSTALCODE",
	}
	for i := 0; i < len(r); i++ {
		rdn := r[len(r)-1-i]
		if i > 0 {
			s += ","
		}
		for j, tv := range rdn {
			if j > 0 {
				s += "+"
			}
			typeString := tv.Type.String()
			typeName, ok := attributeTypeNames[typeString]
			if !ok {
				derBytes, err := asn1.Marshal(tv.Value)
				if err == nil {
					s += typeString + "=#" + hex.EncodeToString(derBytes)
					continue // No value escaping necessary.
				}
				typeName = typeString
			}
			valueString := fmt.Sprint(tv.Value)
			escaped := ""
			begin := 0
			for idx, c := range valueString {
				if (idx == 0 && (c == ' ' || c == '#')) ||
					(idx == len(valueString)-1 && c == ' ') {
					escaped += valueString[begin:idx]
					escaped += "\\" + string(c)
					begin = idx + 1
					continue
				}
				switch c {
				case ',', '+', '"', '\\', '<', '>', ';':
					escaped += valueString[begin:idx]
					escaped += "\\" + string(c)
					begin = idx + 1
				}
			}
			escaped += valueString[begin:]
			s += typeName + "=" + escaped
		}
	}
	return s
}

// getMspIDnID - Get MSP ID and Certificate ID from the certificate of the current caller
func getMspIDnID(stub shim.ChaincodeStubInterface) (string, string, error) {
	// //Check devorgId Tcert attribute
	// // GetCreator returns marshaled serialized identity of the client
	//serializedId, _ := stub.GetCreator()
	id, err := cid.GetID(stub)

	if err != nil {
		return "", "", err
	}

	mspid, err := cid.GetMSPID(stub)

	if err != nil {
		return "", "", err
	}

	//TODO: Delete me
	//logger.Debugf("----------DEBUG---------")
	//logger.Debugf(string(serializedId))
	// logger.Debugf(" User's cert ID: " + id)
	// logger.Debugf(" User's cert ID: " + mspid)

	return mspid, id, nil
}