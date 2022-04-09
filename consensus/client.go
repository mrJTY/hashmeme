package consensus

import (
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
    "encoding/json"
    	"crypto/sha256"
)

type HashMemeMessage struct {
    Author string
    TextContent string
    ImageSha256 string
}

func CreateClient() *hedera.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load environment variables from .env file. Error:\n%s\n", err)
	}

	//Get the operator ID and operator key
	OPERATOR_ID := os.Getenv("OPERATOR_ID")
	OPERATOR_KEY := os.Getenv("OPERATOR_KEY")

	operatorAccountID, err := hedera.AccountIDFromString(OPERATOR_ID)
	if err != nil {
		panic(err)
	}

	operatorKey, err := hedera.PrivateKeyFromString(OPERATOR_KEY)
	if err != nil {
		panic(err)
	}

	client := hedera.ClientForTestnet()
	client.SetOperator(operatorAccountID, operatorKey)

	return client
}

func NewMessage(author string, textContent string) string {
    var hashMemeMessage HashMemeMessage
    hashMemeMessage = HashMemeMessage{Author: author, TextContent: textContent, ImageSha256: "foo"}
    bytes, err := json.Marshal(hashMemeMessage)
    if err != nil {
        log.Fatalf("Could not marshall message %v", err)
    }
    return string(bytes)
}

func HashImage() {
    hash := sha256.Sum256(data)
    return hash
}