package main

import (
  "net/smtp"
//  "os"
)

func main() {
  address := "smtp.gmail.com:587"
  auth := smtp.PlainAuth(
    "",
    "arthurhwoodhouse@gmail.com",
    "coarsesand",
    "smtp.gmail.com",
  )
  from := "arthurhwoodhouse@gmail.com"
  to := []string{"aatina.punjabi@oracle.com", "lloyd.ollerhead@oracle.com"}
  msg := []byte("I'm afraid the lemur got into the pudding cups.")
  if err := smtp.SendMail(address, auth, from, to, msg); err != nil {
    panic(err)
  }
}
