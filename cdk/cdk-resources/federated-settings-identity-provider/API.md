# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnFederatedSettingsIdentityProvider <a name="CfnFederatedSettingsIdentityProvider" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider"></a>

A CloudFormation `MongoDB::Atlas::FederatedSettingsIdentityProvider`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.Initializer"></a>

```typescript
import { CfnFederatedSettingsIdentityProvider } from '@mongodbatlas-awscdk/federated-settings-identity-provider'

new CfnFederatedSettingsIdentityProvider(scope: Construct, id: string, props: CfnFederatedSettingsIdentityProviderProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps">CfnFederatedSettingsIdentityProviderProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps">CfnFederatedSettingsIdentityProviderProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addMetadata"></a>

```typescript
public addMetadata(key: string, value: any): void
```

Add a value to the CloudFormation Resource Metadata.

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.)

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addOverride"></a>

```typescript
public addOverride(path: string, value: any): void
```

Adds an override to the synthesized CloudFormation resource.

To add a
property override, either use `addPropertyOverride` or prefix `path` with
"Properties." (i.e. `Properties.TopicName`).

If the override is nested, separate each nested level using a dot (.) in the path parameter.
If there is an array as part of the nesting, specify the index in the path.

To include a literal `.` in the property name, prefix with a `\`. In most
programming languages you will need to write this as `"\\."` because the
`\` itself will need to be escaped.

For example,
```typescript
cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
```
would add the overrides
```json
"Properties": {
   "GlobalSecondaryIndexes": [
     {
       "Projection": {
         "NonKeyAttributes": [ "myattribute" ]
         ...
       }
       ...
     },
     {
       "ProjectionType": "INCLUDE"
       ...
     },
   ]
   ...
}
```

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.applyRemovalPolicy"></a>

```typescript
public applyRemovalPolicy(policy?: RemovalPolicy, options?: RemovalPolicyOptions): void
```

Sets the deletion policy of the resource based on the removal policy specified.

The Removal Policy controls what happens to this resource when it stops
being managed by CloudFormation, either because you've removed it from the
CDK application or because you've made a change that requires the resource
to be replaced.

The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.getMetadata"></a>

```typescript
public getMetadata(key: string): any
```

Retrieve a value value from the CloudFormation Resource Metadata.

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.)

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.isConstruct"></a>

```typescript
import { CfnFederatedSettingsIdentityProvider } from '@mongodbatlas-awscdk/federated-settings-identity-provider'

CfnFederatedSettingsIdentityProvider.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.isCfnElement"></a>

```typescript
import { CfnFederatedSettingsIdentityProvider } from '@mongodbatlas-awscdk/federated-settings-identity-provider'

CfnFederatedSettingsIdentityProvider.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.isCfnResource"></a>

```typescript
import { CfnFederatedSettingsIdentityProvider } from '@mongodbatlas-awscdk/federated-settings-identity-provider'

CfnFederatedSettingsIdentityProvider.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.attrFederationSettingsId">attrFederationSettingsId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::FederatedSettingsIdentityProvider.FederationSettingsId`. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.attrOktaIdpId">attrOktaIdpId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::FederatedSettingsIdentityProvider.OktaIdpId`. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps">CfnFederatedSettingsIdentityProviderProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrFederationSettingsId`<sup>Required</sup> <a name="attrFederationSettingsId" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.attrFederationSettingsId"></a>

```typescript
public readonly attrFederationSettingsId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::FederatedSettingsIdentityProvider.FederationSettingsId`.

---

##### `attrOktaIdpId`<sup>Required</sup> <a name="attrOktaIdpId" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.attrOktaIdpId"></a>

```typescript
public readonly attrOktaIdpId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::FederatedSettingsIdentityProvider.OktaIdpId`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.props"></a>

```typescript
public readonly props: CfnFederatedSettingsIdentityProviderProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps">CfnFederatedSettingsIdentityProviderProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProvider.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/federated-settings-identity-provider.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/federated-settings-identity-provider.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/federated-settings-identity-provider'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/federated-settings-identity-provider.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/federated-settings-identity-provider.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnFederatedSettingsIdentityProviderProps <a name="CfnFederatedSettingsIdentityProviderProps" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps"></a>

Returns, adds, edits, and removes federation-related features such as role mappings and connected organization configurations.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.Initializer"></a>

```typescript
import { CfnFederatedSettingsIdentityProviderProps } from '@mongodbatlas-awscdk/federated-settings-identity-provider'

const cfnFederatedSettingsIdentityProviderProps: CfnFederatedSettingsIdentityProviderProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.associatedDomains">associatedDomains</a></code> | <code>string[]</code> | List that contains the domains associated with the identity provider. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.displayName">displayName</a></code> | <code>string</code> | Human-readable label that identifies the IdP. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.issuerUri">issuerUri</a></code> | <code>string</code> | Unique string that identifies the issuer of the SAML Assertion. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.requestBinding">requestBinding</a></code> | <code>string</code> | SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.responseSignatureAlgorithm">responseSignatureAlgorithm</a></code> | <code>string</code> | Signature algorithm that Federated Authentication uses to encrypt the identity provider signature. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.ssoDebugEnabled">ssoDebugEnabled</a></code> | <code>boolean</code> | Flag that indicates whether the identity provider has SSO debug enabled. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.ssoUrl">ssoUrl</a></code> | <code>string</code> | URL that points to the receiver of the SAML authentication request. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.status">status</a></code> | <code>string</code> | String enum that indicates whether the identity provider is active. |

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-settings-identity-provider.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `associatedDomains`<sup>Optional</sup> <a name="associatedDomains" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.associatedDomains"></a>

```typescript
public readonly associatedDomains: string[];
```

- *Type:* string[]

List that contains the domains associated with the identity provider.

---

##### `displayName`<sup>Optional</sup> <a name="displayName" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.displayName"></a>

```typescript
public readonly displayName: string;
```

- *Type:* string

Human-readable label that identifies the IdP.

---

##### `issuerUri`<sup>Optional</sup> <a name="issuerUri" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.issuerUri"></a>

```typescript
public readonly issuerUri: string;
```

- *Type:* string

Unique string that identifies the issuer of the SAML Assertion.

---

##### `requestBinding`<sup>Optional</sup> <a name="requestBinding" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.requestBinding"></a>

```typescript
public readonly requestBinding: string;
```

- *Type:* string

SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request.

---

##### `responseSignatureAlgorithm`<sup>Optional</sup> <a name="responseSignatureAlgorithm" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.responseSignatureAlgorithm"></a>

```typescript
public readonly responseSignatureAlgorithm: string;
```

- *Type:* string

Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.

---

##### `ssoDebugEnabled`<sup>Optional</sup> <a name="ssoDebugEnabled" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.ssoDebugEnabled"></a>

```typescript
public readonly ssoDebugEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether the identity provider has SSO debug enabled.

---

##### `ssoUrl`<sup>Optional</sup> <a name="ssoUrl" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.ssoUrl"></a>

```typescript
public readonly ssoUrl: string;
```

- *Type:* string

URL that points to the receiver of the SAML authentication request.

---

##### `status`<sup>Optional</sup> <a name="status" id="@mongodbatlas-awscdk/federated-settings-identity-provider.CfnFederatedSettingsIdentityProviderProps.property.status"></a>

```typescript
public readonly status: string;
```

- *Type:* string

String enum that indicates whether the identity provider is active.

---


