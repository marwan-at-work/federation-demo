// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package reviews

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql/introspection"
)

func (ec *executionContext) __resolve__service(ctx context.Context) (introspection.Service, error) {
	if ec.DisableIntrospection {
		return introspection.Service{}, errors.New("federated introspection disabled")
	}
	return introspection.Service{
		SDL: `type Product @extends @key(fields: "upc") {
	upc: String! @external
	reviews: [Review]
}
type Review @key(fields: "id") {
	id: ID!
	body: String
	author: User @provides(fields: "username")
	product: Product
}
type User @extends @key(fields: "id") {
	id: ID! @external
	username: String @external
	reviews: [Review]
}
`,
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) ([]_Entity, error) {
	list := []_Entity{}
	for _, rep := range representations {
		typeName, ok := rep["__typename"].(string)
		if !ok {
			return nil, errors.New("__typename must be an existing string")
		}
		switch typeName {

		case "User":
			id, ok := rep["id"].(string)
			if !ok {
				return nil, errors.New("opsies")
			}
			resp, err := ec.resolvers.Entity().FindUserByID(ctx, id)
			if err != nil {
				return nil, err
			}

			list = append(list, resp)

		case "Product":
			id, ok := rep["upc"].(string)
			if !ok {
				return nil, errors.New("opsies")
			}
			resp, err := ec.resolvers.Entity().FindProductByUpc(ctx, id)
			if err != nil {
				return nil, err
			}

			list = append(list, resp)

		case "Review":
			id, ok := rep["id"].(string)
			if !ok {
				return nil, errors.New("opsies")
			}
			resp, err := ec.resolvers.Entity().FindReviewByID(ctx, id)
			if err != nil {
				return nil, err
			}

			list = append(list, resp)

		default:
			return nil, errors.New("unknown type: " + typeName)
		}
	}
	return list, nil
}
