package document

import "github.com/jensneuse/graphql-go-tools/pkg/lexing/position"

// TypeSystemDefinition as specified in:
// http://facebook.github.io/graphql/draft/#TypeSystemDefinition
type TypeSystemDefinition struct {
	SchemaDefinition           SchemaDefinition
	ScalarTypeDefinitions      []int
	ObjectTypeDefinitions      []int
	InterfaceTypeDefinitions   []int
	UnionTypeDefinitions       []int
	EnumTypeDefinitions        []int
	InputObjectTypeDefinitions []int
	DirectiveDefinitions       []int
	Position                   position.Position
}

func (t TypeSystemDefinition) NodeSelectionSet() int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeInputFieldsDefinition() int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeInputValueDefinitions() []int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodePosition() position.Position {
	return t.Position
}

func (t TypeSystemDefinition) NodeValueType() ValueType {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeValueReference() int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeUnionMemberTypes() []int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeName() int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeAlias() int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeDescription() ByteSliceReference {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeArgumentSet() int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeArgumentsDefinition() int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeDirectiveSet() int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeEnumValuesDefinition() []int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeFields() []int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeFieldsDefinition() []int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeFragmentSpreads() []int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeInlineFragments() []int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeVariableDefinitions() []int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeType() int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeOperationType() OperationType {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeValue() int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeDefaultValue() int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeImplementsInterfaces() []int {
	panic("implement me")
}

func (t TypeSystemDefinition) NodeSchemaDefinition() SchemaDefinition {
	return t.SchemaDefinition
}

func (t TypeSystemDefinition) NodeScalarTypeDefinitions() []int {
	return t.ScalarTypeDefinitions
}

func (t TypeSystemDefinition) NodeObjectTypeDefinitions() []int {
	return t.ObjectTypeDefinitions
}

func (t TypeSystemDefinition) NodeInterfaceTypeDefinitions() []int {
	return t.InterfaceTypeDefinitions
}

func (t TypeSystemDefinition) NodeUnionTypeDefinitions() []int {
	return t.UnionTypeDefinitions
}

func (t TypeSystemDefinition) NodeEnumTypeDefinitions() []int {
	return t.EnumTypeDefinitions
}

func (t TypeSystemDefinition) NodeInputObjectTypeDefinitions() []int {
	return t.InputObjectTypeDefinitions
}

func (t TypeSystemDefinition) NodeDirectiveDefinitions() []int {
	return t.DirectiveDefinitions
}

type TypeSystemDefinitions []TypeSystemDefinition
