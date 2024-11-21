package graph

import (
	"auth-system/graph/model"
	"github.com/graphql-go/graphql"
)

func NewSchema() *graphql.Schema {
	// Define types
	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"email":     &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"firstName": &graphql.Field{Type: graphql.String},
			"lastName":  &graphql.Field{Type: graphql.String},
			"createdAt": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"updatedAt": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	})

	authResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "AuthResponse",
		Fields: graphql.Fields{
			"token": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"user":  &graphql.Field{Type: graphql.NewNonNull(userType)},
		},
	})

	// Define input types
	registerInput := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "RegisterInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"email":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"password":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"firstName": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"lastName":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		},
	})

	loginInput := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "LoginInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"email":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"password": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		},
	})

	resolver := NewResolver()

	// Define root Query type
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.Query().Me(p.Context)
				},
			},
		},
	})

	// Define root Mutation type
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"register": &graphql.Field{
				Type: authResponseType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{Type: registerInput},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})
					registerInput := model.RegisterInput{
						Email:     input["email"].(string),
						Password:  input["password"].(string),
						FirstName: input["firstName"].(string),
						LastName:  input["lastName"].(string),
					}
					return resolver.Mutation().Register(p.Context, registerInput)
				},
			},
			"login": &graphql.Field{
				Type: authResponseType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{Type: loginInput},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})
					loginInput := model.LoginInput{
						Email:    input["email"].(string),
						Password: input["password"].(string),
					}
					return resolver.Mutation().Login(p.Context, loginInput)
				},
			},
		},
	})

	// Create Schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		panic(err)
	}

	return &schema
}
