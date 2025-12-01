package base

// IDataMapper is the base interface for all data mappers
type IDataMapper interface{}

// BaseDataMapper provides generic mapping functionality between Domain, Entity (DB Model), and DTO
type BaseDataMapper[TDomain any, TEntity any, TDto any] struct {
	// ToDomain converts a database entity to a domain model
	ToDomainFunc func(entity TEntity) TDomain
	// ToDomains converts multiple database entities to domain models
	ToDomainsFunc func(entities []TEntity) []TDomain
	// ToEntity converts a domain model to a database entity
	ToEntityFunc func(domain TDomain) TEntity
	// ToDTO converts a domain model to a DTO
	ToDTOFunc func(domain TDomain) TDto
	// ToDTOs converts multiple domain models to DTOs
	ToDTOsFunc func(domains []TDomain) []TDto
}

// NewBaseDataMapper creates a new base data mapper with default implementations
func NewBaseDataMapper[TDomain any, TEntity any, TDto any](
	toDomain func(TEntity) TDomain,
	toEntity func(TDomain) TEntity,
	toDTO func(TDomain) TDto,
) *BaseDataMapper[TDomain, TEntity, TDto] {
	mapper := &BaseDataMapper[TDomain, TEntity, TDto]{
		ToDomainFunc: toDomain,
		ToEntityFunc: toEntity,
		ToDTOFunc:    toDTO,
	}

	// Default implementation for batch conversions
	mapper.ToDomainsFunc = func(entities []TEntity) []TDomain {
		domains := make([]TDomain, len(entities))
		for i, entity := range entities {
			domains[i] = mapper.ToDomainFunc(entity)
		}
		return domains
	}

	mapper.ToDTOsFunc = func(domains []TDomain) []TDto {
		dtos := make([]TDto, len(domains))
		for i, domain := range domains {
			dtos[i] = mapper.ToDTOFunc(domain)
		}
		return dtos
	}

	return mapper
}

// ToDomain converts entity to domain
func (m *BaseDataMapper[TDomain, TEntity, TDto]) ToDomain(entity TEntity) TDomain {
	return m.ToDomainFunc(entity)
}

// ToDomains converts entities to domains
func (m *BaseDataMapper[TDomain, TEntity, TDto]) ToDomains(entities []TEntity) []TDomain {
	return m.ToDomainsFunc(entities)
}

// ToEntity converts domain to entity
func (m *BaseDataMapper[TDomain, TEntity, TDto]) ToEntity(domain TDomain) TEntity {
	return m.ToEntityFunc(domain)
}

// ToDTO converts domain to DTO
func (m *BaseDataMapper[TDomain, TEntity, TDto]) ToDTO(domain TDomain) TDto {
	return m.ToDTOFunc(domain)
}

// ToDTOs converts domains to DTOs
func (m *BaseDataMapper[TDomain, TEntity, TDto]) ToDTOs(domains []TDomain) []TDto {
	return m.ToDTOsFunc(domains)
}
