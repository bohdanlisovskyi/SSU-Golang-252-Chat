package contacts

var ContactsList *Contacts

type Contact struct {
	UserName       string
	NickName       string
	BirthDay       int
	MessageHistory string
	IsBlocked      bool
}

type Contacts struct {
	ContactsList   []Contact
	CurrentContact int
}

//IndexByUserName return index of contact with userName if it exist, or -1 otherwise
func (cont *Contacts) IndexByUserName(userName string) int {
	for index, contact := range cont.ContactsList {
		if contact.UserName == userName {
			return index
		}
	}
	return -1
}

func (cont *Contacts) GetContactByUserName(userName string) *Contact {
	for _, contact := range cont.ContactsList {
		if contact.UserName == userName {
			return &contact
		}
	}
	return nil
}
