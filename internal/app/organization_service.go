package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type OrganizationService interface {
	Save(o domain.Organization) (domain.Organization, error)
}

type organizationService struct {
	organizationRepo database.OrganizationRepository
}

func NewOrganizationService(or database.OrganizationRepository) OrganizationService {
	return organizationService{
		organizationRepo: or,
	}
}

func (s organizationService) Save(o domain.Organization) (domain.Organization, error) {
	o, err := s.organizationRepo.Save(o)
	if err != nil {
		log.Printf("OrganizationService: %s", err)
		return domain.Organization{}, err
	}

	return o, nil
}
