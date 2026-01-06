package access

import "encoding/json"

type UserIdentity struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Message struct {
	Sender   UserIdentity `json:"sender"`
	Receiver UserIdentity `json:"receiver"`
	Message  string       `json:"message"`
}

func stringToByte(s string) []byte {
	return []byte(s)
}

func messageToBytes(
	senderUsername, senderEmail string,
	receiverUsername, receiverEmail string,
	message string,
) ([]byte, error) {

	m := Message{
		Sender: UserIdentity{
			Username: senderUsername,
			Email:    senderEmail,
		},
		Receiver: UserIdentity{
			Username: receiverUsername,
			Email:    receiverEmail,
		},
		Message: message,
	}

	return json.Marshal(m)
}
