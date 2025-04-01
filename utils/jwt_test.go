package utils

import (
	"fmt"
	"testing"
)

func TestGenerateJwtToken(t *testing.T) {
	jwtToken, err := GenerateJwtToken(10, "Mundo", "10086")
	if err != nil {
		t.Errorf("call GenerateJwtToken failed, error: %v", err)
	}
	fmt.Println(jwtToken)
}

func TestParseJwtToken(t *testing.T) {
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" +
		".eyJpc3MiOiJNdW5kbyIsImV4cCI6MTc0MzczMzY4NCwiaWF0IjoxNzQzNDc0NDg0LCJqdGkiOiIyOTU5OGRiYi03NmVhLTQ5NjMtYjA2ZS1lNzRhNDU3ZGY1NmMiLCJ1c2VySWQiOjEwLCJ1c2VyTmFtZSI6Ik11bmRvIiwicGhvbmUiOiIxMDA4NiJ9" +
		".vG_7vtL4hbjTQYF5nnvm1xDDijkCXSXCqd9d8KNF4Dg"
	claims, err := ParseJwtToken(jwtToken)
	if err != nil {
		t.Errorf("call ParseJwtToken failed, error: %v", err)
	}
	fmt.Printf("UserID: %d, UserName: %s, Phone: %s\n", claims.UserID, claims.UserName, claims.Phone)
}
