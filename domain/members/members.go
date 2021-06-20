package members

func GetAll() ([]Member, error) {
	return getAll()
}

func PutMember(member Member) error {
	if err := validateMember(member); err != nil {
		return err
	}

	if member.ID != 0 {
		return updateMember(member)
	}
	return addMember(member)
}

func DeleteMember(id int) error {
	return deleteMember(id)
}
