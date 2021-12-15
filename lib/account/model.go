package account

import "gorm.io/gorm"

type Model struct {
	db *gorm.DB
}

func NewModel(db *gorm.DB) Model {
	return Model{db}
}

func (m Model) Create(input CreateAccountInput) (Account, error) {
	account := Account{ClientId: input.ClientId, Amount: input.Amount}
	err := m.db.Create(&account).Error
	return account, err
}

func (m Model) UpdateById(accountId uint, input UpdateAccountInput) (Account, error) {
	var account Account
	if err := m.FindById(&account, accountId); err != nil {
		return account, err
	}
	err := m.db.Model(&account).Updates(input).Error
	return account, err
}

func (m Model) Update(account *Account, input UpdateAccountInput) error {
	return m.db.Model(account).Where("id", input.Id).Updates(input).Error
}

func (m Model) DeleteById(accountId uint) error {
	var account Account
	m.FindById(&account, accountId)
	return m.Delete(account)
}

func (m Model) Delete(account Account) error {
	return m.db.Delete(&account, account.Id).Error
}

func (m Model) FindAllLimit(accounts *[]Account, limit int) error {
	return m.db.Limit(limit).Find(accounts).Error
}

func (m Model) FindById(account *Account, accountId uint) error {
	return m.db.First(&account, accountId).Error
}

func (m Model) FindByIds(accounts *[]Account, accountIds []uint, limit int) error {
	return m.db.Limit(limit).Find(accounts, accountIds).Error
}
