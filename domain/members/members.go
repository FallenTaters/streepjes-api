package members

func GetAll() ([]Member, error) {
	return repo.getAll()
}

func Put(member Member) error {
	if err := validateMember(member); err != nil {
		return err
	}

	if member.ID != 0 {
		return repo.updateMember(member)
	}
	return repo.addMember(member)
}

func ForceDelete(id int) error {
	return repo.deleteMember(id)
}
