package main

import (
	"testing"
	"time"

	"github.com/jncmaguire/example-proto-custom-resolver/pet"
	taxonomy "github.com/jncmaguire/example-proto-types"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestProcessMessageAsMap(t *testing.T) {
	scenarios := []struct {
		name         string
		givenMessage proto.Message
		expectedMap  map[string]interface{}
	}{
		{
			name: "empty",
		},
		{
			name: "pet",
			givenMessage: &pet.Pet{
				Name:        "PatchyScratchy",
				DateOfBirth: timestamppb.New(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)),
				Animal: &pet.Animal{
					Classification: &taxonomy.Classification{
						Kingdom:   "Animalia",
						Phylum:    "Chordata",
						Class:     "Mammalia",
						Order:     "Carnivora",
						Suborder:  "Feliformia",
						Family:    "Felidae",
						Subfamily: "Felinae",
						Genus:     "Felis",
						Species:   "F.catus",
					},
					Names: []string{
						"cat",
						"kitty",
					},
				},
				Friends: []*pet.Animal{
					&pet.Animal{
						Classification: &taxonomy.Classification{
							Kingdom:   "Animalia",
							Phylum:    "Chordata",
							Class:     "Mammalia",
							Order:     "Primates",
							Suborder:  "Haplorhini",
							Family:    "Hominidae",
							Subfamily: "Homininae",
							Genus:     "Homo",
							Species:   "H. sapiens",
						},
						Names: []string{
							"man",
						},
					},
				},
				Parents: map[string]*pet.Animal{
					"cat-mom": &pet.Animal{
						Classification: &taxonomy.Classification{
							Kingdom:   "Animalia",
							Phylum:    "Chordata",
							Class:     "Mammalia",
							Order:     "Primates",
							Suborder:  "Haplorhini",
							Family:    "Hominidae",
							Subfamily: "Homininae",
							Genus:     "Homo",
							Species:   "H. sapiens",
						},
						Names: []string{
							"man",
						},
					},
					"human-mom": &pet.Animal{
						Classification: &taxonomy.Classification{
							Kingdom:   "Animalia",
							Phylum:    "Chordata",
							Class:     "Mammalia",
							Order:     "Carnivora",
							Suborder:  "Feliformia",
							Family:    "Felidae",
							Subfamily: "Felinae",
							Genus:     "Felis",
							Species:   "F.catus",
						},
						Names: []string{
							"cat",
							"kitty",
						},
					},
				},
			},
			expectedMap: map[string]interface{}{
				"name":        "PatchyScratchy",
				"dateOfBirth": timestamppb.New(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)),
				"animal": &pet.Animal{
					Classification: &taxonomy.Classification{
						Kingdom:   "Animalia",
						Phylum:    "Chordata",
						Class:     "Mammalia",
						Order:     "Carnivora",
						Suborder:  "Feliformia",
						Family:    "Felidae",
						Subfamily: "Felinae",
						Genus:     "Felis",
						Species:   "F.catus",
					},
					Names: []string{
						"cat",
						"kitty",
					},
				},
				"friends": []interface{}{
					&pet.Animal{
						Classification: &taxonomy.Classification{
							Kingdom:   "Animalia",
							Phylum:    "Chordata",
							Class:     "Mammalia",
							Order:     "Primates",
							Suborder:  "Haplorhini",
							Family:    "Hominidae",
							Subfamily: "Homininae",
							Genus:     "Homo",
							Species:   "H. sapiens",
						},
						Names: []string{
							"man",
						},
					},
				},
				"parents": map[string]interface{}{
					"cat-mom": &pet.Animal{
						Classification: &taxonomy.Classification{
							Kingdom:   "Animalia",
							Phylum:    "Chordata",
							Class:     "Mammalia",
							Order:     "Primates",
							Suborder:  "Haplorhini",
							Family:    "Hominidae",
							Subfamily: "Homininae",
							Genus:     "Homo",
							Species:   "H. sapiens",
						},
						Names: []string{
							"man",
						},
					},
					"human-mom": &pet.Animal{
						Classification: &taxonomy.Classification{
							Kingdom:   "Animalia",
							Phylum:    "Chordata",
							Class:     "Mammalia",
							Order:     "Carnivora",
							Suborder:  "Feliformia",
							Family:    "Felidae",
							Subfamily: "Felinae",
							Genus:     "Felis",
							Species:   "F.catus",
						},
						Names: []string{
							"cat",
							"kitty",
						},
					},
				},
			},
		},
		{
			name:         "no-add",
			givenMessage: &pet.Blank{},
			expectedMap:  map[string]interface{}{},
		},
	}

	for i := range scenarios {
		t.Run(scenarios[i].name, func(t *testing.T) {
			actual := processMessageAsMap(scenarios[i].givenMessage)

			assertEqualContainers(t, scenarios[i].expectedMap, actual)
		})

	}
}

func assertEqualObjects(t *testing.T, expected interface{}, actual interface{}, key interface{}) {
	switch expectedValue := expected.(type) {
	case proto.Message:
		actualValue, ok := actual.(proto.Message)
		assert.True(t, ok)
		assert.Truef(t, proto.Equal(expectedValue, actualValue), "%v values should match", key)
	default:
		assert.Equalf(t, expected, actual, "%v values should match", key)
	}
}

func assertEqualContainers(t *testing.T, expected map[string]interface{}, actual map[string]interface{}) {
	assert.Len(t, actual, len(expected))
	for key, _ := range expected {
		switch expectedValue := expected[key].(type) {
		case []interface{}:
			actualValue, ok := actual[key].([]interface{})
			assert.True(t, ok)
			assert.Len(t, actualValue, len(expectedValue))
			for i := range expectedValue {
				assertEqualObjects(t, expectedValue[i], actualValue[i], key)
			}
		case map[string]interface{}:
			actualValue, ok := actual[key].(map[string]interface{})
			assert.True(t, ok)
			assertEqualContainers(t, expectedValue, actualValue)
		case map[bool]interface{}:
			actualValue, ok := actual[key].(map[bool]interface{})
			assert.True(t, ok)
			assert.Len(t, actualValue, len(expectedValue))
			for k, _ := range expectedValue {
				assertEqualObjects(t, expectedValue[k], actualValue[k], k)
			}
		case map[int64]interface{}:
			actualValue, ok := actual[key].(map[int64]interface{})
			assert.True(t, ok)
			assert.Len(t, actualValue, len(expectedValue))
			for k, _ := range expectedValue {
				assertEqualObjects(t, expectedValue[k], actualValue[k], k)
			}
		case map[uint64]interface{}:
			actualValue, ok := actual[key].(map[uint64]interface{})
			assert.True(t, ok)
			assert.Len(t, actualValue, len(expectedValue))
			for k, _ := range expectedValue {
				assertEqualObjects(t, expectedValue[k], actualValue[k], k)
			}
		default:
			assertEqualObjects(t, expected[key], actual[key], key)
		}
	}
}
