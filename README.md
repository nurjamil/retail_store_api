# Simple Inventory API
This is a simple API

## Version: 1.2

**Contact information:**  
you@your-company.com  

**License:** [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### /items

#### GET
##### Summary

item list

##### Description

By passing in the appropriate options, you can search for
certain available item

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| items_id | query | pass an optional search string for looking up item | No | string |
| category | query | pass an optional search string for looking up item | No | string |
| name | query | pass an optional search string for looking up item | No | string |
| limit | query | maximum number of records to return | No | integer (uint32) |
| minPrice | query | minimum price | No | integer (uint32) |
| maxPrice | query | maximum price | No | integer (uint32) |
| sort | query | pass an optional search string for looking up item | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | search results matching criteria | object |
| 400 | bad input parameter | object |

#### POST
##### Summary

adds an inventory item

##### Description

Adds an item to the system

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| inventoryItem | body | Inventory item to add | No | [Item](#item) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | item created | [Response](#response) |
| 400 | invalid input, object invalid | [Response](#response) |
| 409 | --- | [Response](#response) |

#### PUT
##### Summary

edit an inventory item

##### Description

edit an item to the system

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| inventoryItem | body | Inventory item to add | No | [Item](#item) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | item updated | [Response](#response) |
| 400 | invalid input, object invalid | [Response](#response) |

### /items/{item_name}

#### GET
##### Summary

item list

##### Description

By passing in the appropriate options, you can search for
certain available item

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| item_name | path | item name user to get | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | search results matching criteria | object |
| 400 | bad input parameter | object |

### /shopping_carts

#### GET
##### Summary

shopping cart list

##### Description

By passing in the appropriate options, you can get items in shopping cart

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | pass an required search string for looking up item in shopping cart | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | ok results matching criteria | [Response](#response) |
| 400 | bad input parameter | [Response](#response) |

#### POST
##### Summary

adds an item to shopping cart

##### Description

Adds an item to the system

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| authorization | header | token | Yes | string |
| item_id | body | Inventory item to add | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | item created | [Response](#response) |
| 400 | invalid input, object invalid | [Response](#response) |
| 409 | an existing item already exists | [Response](#response) |

#### DELETE
##### Summary

delete an item from shopping cart

##### Description

delete an item from the system

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| authorization | header | token | Yes | string |
| item_id | body | Inventory item to add | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | item deleted | [Response](#response) |
| 400 | invalid input, object invalid | [Response](#response) |
| 409 | item not exist | [Response](#response) |

### /orders/{order_id}

#### GET
##### Summary

shopping cart list

##### Description

By passing in the appropriate options, you can get items in shopping cart

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| order_id | path |  | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | search results matching criteria | object |
| 400 | bad input parameter | object |

### /orders

#### GET
##### Summary

shopping cart list

##### Description

By passing in the appropriate options, you can get items in shopping cart

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | pass an required search string for looking up ite  in shopping cart | Yes | string |
| order_id | query | pass an order_id | No | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | search results matching criteria | object |
| 400 | bad input parameter | object |

#### POST
##### Summary

adds an order

##### Description

Adds an order to the system

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| authorization | header | token | Yes | string |
| item | body | Inventory item to add | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | order created | [Response](#response) |
| 400 | invalid input, object invalid | [Response](#response) |

#### DELETE
##### Summary

delete order item

##### Description

delete an order from the system

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| authorization | header | token | Yes | string |
| order_id | body | Inventory item to add | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | order deleted | [Response](#response) |
| 400 | invalid input, object invalid | [Response](#response) |
| 409 | order not exist | [Response](#response) |

### /register

#### POST
##### Summary

register user account

##### Description

register user account to the system

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| register user | body | register user account | Yes | [RegisterUser](#registeruser) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | account created | object |
| 400 | invalid input, object invalid | object |

### /login

#### POST
##### Summary

login user account

##### Description

login user account to the system

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| user login | body | login user account | Yes | [UserLogin](#userlogin) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | login access granted | object |
| 400 | incorrect username/password | object |

### /payments

#### GET
##### Summary

payment list

##### Description

By passing in the appropriate options, you can search for
certain available payment

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| authorization | header |  | Yes | string |
| payment_id | query | pass an optional search string for looking up item | No | string |
| order_id | query | order id from order | No | string |
| customer_id | query | customer id from customer | No | string |
| status | query | status | No | integer (uint32) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | search results matching criteria | object |
| 400 | bad input parameter | object |

### /payments/{payment_id}

#### GET
##### Summary

payment list

##### Description

By passing in the appropriate options, you can search for
certain available payment

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| authorization | header |  | Yes | string |
| payment_id | path | pass an optional search integer for looking up payment | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | search results matching criteria | object |
| 400 | bad input parameter | object |

### /courier

#### GET
##### Summary

courier list

##### Description

By passing in the appropriate options, you can search for
certain available courier

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| authorization | header |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | search results matching criteria | [ [Courier](#courier) ] |
| 400 | bad input parameter |  |

### /address

#### GET
##### Summary

get address

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| authorization | header |  | Yes | string |
| address_id | query | pass an optional search string for looking up address | No | integer |
| name | query | pass an optional search string for looking up address | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | search results matching criteria | object |
| 400 | bad input parameter | object |

#### POST
##### Summary

adds an address item

##### Description

Adds an address to the system

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| authorization | header | authorization token | Yes | string |
| addressItem | body | address item to add | Yes | object |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | address created | object |
| 400 | invalid input, object invalid | object |
| 409 | an existing address already exists | object |

### /address/{address_id}

#### GET
##### Summary

get address

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| authorization | header |  | Yes | string |
| address_id | path | pass an optional search string for looking up address | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | search results matching criteria | object |
| 400 | bad input parameter | object |

### Models

#### Item

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | integer (uint32) | _Example:_ `69` | No |
| name | string | _Example:_ `"Widget Adapter"` | No |
| description | string | _Example:_ `"ini adalah buku budi"` | No |
| category | string |  | No |
| price | integer (uint32) |  | No |

#### shopping_cart

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| shopping_cart | array |  |  |

#### order

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| order_id | integer |_Example:_ `1` | No |
| courier | object |  | No |
| total_amount | integer |  | No |
| address | string |  | No |
| payment_id | string |  | No |
| status | string |  | No |

#### RegisterUser

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| username | string |  | No |
| nama | string |  | No |
| email | string | _Example:_ `"email@email.com"` | No |
| phone_number | string |_Example:_ `"+628999"` | No |
| password | string | _Example:_ `"password123"` | No |

#### UserLogin

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| username | string |  | No |
| password | string | _Example:_ `"password123"` | No |

#### Payment

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| payment_id | integer |  | No |
| transfer_code | string |  | No |
| order_id | integer |  | No |
| customer_id | integer |  | No |
| status | string |  | No |
| total_amount | integer |  | No |

#### Courier

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| courier_id | integer |  | No |
| company_name | string |  | No |

#### Address

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| address_id | string |  | No |
| name | string |  | No |
| address | string |  | No |

#### Response

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | integer |  | No |
| status | string |  | No |
| message | string |  | No |

#### CustomerInfo

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| customer_id | integer |  | No |
| username | string |  | No |
| email | string |  | No |
| no_hp | string |  | No |
| token | string |  | No |
