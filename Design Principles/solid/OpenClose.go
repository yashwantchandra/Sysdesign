package main


import "fmt"

// MessageSender is an abstraction.
// Notifier depends on this, not on concrete types.

// What shows Open/Closed here?
// Notifier depends on the MessageSender interface, not concrete types.
// To add a new notification channel (Slack, Push, WhatsApp), you only add a new struct that implements MessageSender.
// Existing, tested code (Notifier, EmailSender, SMSSender) stays unchanged, satisfying the Open/Closed Principle.


type MessageSender interface {
	SendMessage(to, message string) error
}

// EmailSender is one implementation (email).
type EmailSender struct{}

func (e EmailSender) SendMessage(to, message string) error {
	fmt.Printf("Sending EMAIL to %s: %s\n", to, message)
	return nil
}

// SMSSender is another implementation (SMS).
type SMSSender struct{}

func (s SMSSender) SendMessage(to, message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", to, message)
	return nil
}

// Notifier is closed for modification: its code does not change
// when we add new senders. It just uses the MessageSender interface.
type Notifier struct {
	sender MessageSender
}

func NewNotifier(sender MessageSender) *Notifier {
	return &Notifier{sender: sender}
}

func (n *Notifier) NotifyUser(userID, message string) error {
	// In real life, you might look up user contact details here.
	to := userID
	return n.sender.SendMessage(to, message)
}

func main() {
	// Use email
	emailNotifier := NewNotifier(EmailSender{})
	emailNotifier.NotifyUser("alice@example.com", "Welcome Alice!")

	// Use SMS
	smsNotifier := NewNotifier(SMSSender{})
	smsNotifier.NotifyUser("+1234567890", "Your OTP is 1234")

	// In the future, to add Slack notifications:
	// 1. Create SlackSender struct implementing MessageSender.
	// 2. Construct Notifier with SlackSender.
	// No change needed in Notifier itself.
}