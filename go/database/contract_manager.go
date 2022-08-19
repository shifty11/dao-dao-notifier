package database

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/types"
)

type ContractManager struct {
	client *ent.Client
	ctx    context.Context
}

func NewContractManager(client *ent.Client, ctx context.Context) *ContractManager {
	return &ContractManager{client: client, ctx: ctx}
}

type ContractStatus string

const (
	ContractCreated      ContractStatus = "created"
	ContractUpdated      ContractStatus = "updated"
	ContractImageChanged ContractStatus = "image_changed"
	ContractNoChanges    ContractStatus = "no_changes"
)

// CreateOrUpdate creates a new contract or updates an existing one
//
// returns (contract, created)
func (m *ContractManager) CreateOrUpdate(contractAddr string, config *types.ContractData) (*ent.Contract, ContractStatus) {
	c, err := m.client.Contract.
		Query().
		Where(contract.AddressEQ(contractAddr)).
		First(m.ctx)
	if err != nil && ent.IsNotFound(err) {
		c, err = m.client.Contract.
			Create().
			SetAddress(contractAddr).
			SetName(config.Name).
			SetDescription(config.Description).
			SetImageURL(config.ImageUrl).
			Save(m.ctx)
		if err != nil {
			log.Sugar.Panicf("Error while creating contract: %v", err)
		}
		return c, ContractCreated
	} else if err != nil {
		log.Sugar.Panicf("Error while querying contract: %v", err)
	} else {
		if c.Name != config.Name || c.Description != config.Description || c.ImageURL != config.ImageUrl {
			updated, err := m.client.Contract.
				UpdateOne(c).
				SetName(config.Name).
				SetDescription(config.Description).
				SetImageURL(config.ImageUrl).
				Save(m.ctx)
			if err != nil {
				log.Sugar.Panicf("Error while updating contract: %v", err)
			}
			if c.ImageURL != updated.ImageURL {
				return updated, ContractImageChanged
			}
			return updated, ContractUpdated
		}
	}
	return c, ContractNoChanges
}

func (m *ContractManager) All() []*ent.Contract {
	all, err := m.client.Contract.
		Query().
		All(m.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while querying contracts: %v", err)
	}
	return all
}

func (m *ContractManager) Get(id int) (*ent.Contract, error) {
	return m.client.Contract.
		Query().
		Where(contract.ID(id)).
		Only(m.ctx)
}

func (m *ContractManager) SaveThumbnailUrl(entContract *ent.Contract, url string) *ent.Contract {
	updated, err := m.client.Contract.
		UpdateOne(entContract).
		SetThumbnailURL(url).
		Save(m.ctx)
	if err != nil {
		log.Sugar.Panicf("Error while updating contract: %v", err)
	}
	return updated
}
