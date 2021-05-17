
---
# Condition


<details><summary style="margin:3px"> Search conditions</summary>

<p style="margin:10px">
    <span style="background-color:rgb(0, 122, 202); color:white; padding:2px"> 
        GET
    </span>
    <code> /v1/conditions </code>
</p>

---

### Query Parameters: `id` | `io_id` |






### Response
```json

[
    {
        "id": 10733652028434,
        "title": "coo coo",
        "child_trigger_id": 2,
        "io_id": 339
    },
    {
        "id": 10737777178551,
        "title": "okey",
        "child_trigger_id": 2,
        "io_id": 339
    },
    {
        "id": 35941932369724,
        "title": "cond feb",
        "child_trigger_id": 10,
        "io_id": 339
    },
    {
        "id": 43579629179435,
        "title": "cond feb 11",
        "child_trigger_id": 8,
        "io_id": 339
    }
]
```


### Response Fields: 
| Field | Required | Description |
| ------ | ------ | ------ |
|id|false|condition id|
|title|false|condition name|


</details>
---


<details><summary style="margin:3px"> Save condition</summary>

<p style="margin:10px">
    <span style="background-color:rgb(0, 122, 202); color:white; padding:2px"> 
        POST
    </span>
    <code> /v1/conditions </code>
</p>

---



### ***Request Body Schema***: 
```json
{
  "id": 123,
  "name": "Kale"
}
```



### Request Fields: 
| Field | Required | Description |
| ------ | ------ | ------ |
|id|true|the user id|
|name|false|the user name|


### Response
```json
{"status": 201, "message": "created"}
```



</details>
---


---
# RuleSet


<details><summary style="margin:3px"> Search Rulesets</summary>

<p style="margin:10px">
    <span style="background-color:rgb(0, 122, 202); color:white; padding:2px"> 
        GET
    </span>
    <code> /v1/conditions </code>
</p>

---

### Query Parameters: `id` | `io_id` |






### Response
```json

[
    {
        "id": 10733652028434,
        "title": "coo coo",
        "child_trigger_id": 2,
        "io_id": 339
    },
    {
        "id": 10737777178551,
        "title": "okey",
        "child_trigger_id": 2,
        "io_id": 339
    },
    {
        "id": 35941932369724,
        "title": "cond feb",
        "child_trigger_id": 10,
        "io_id": 339
    },
    {
        "id": 43579629179435,
        "title": "cond feb 11",
        "child_trigger_id": 8,
        "io_id": 339
    }
]
```


### Response Fields: 
| Field | Required | Description |
| ------ | ------ | ------ |
|id|false|condition id|
|title|false|condition name|


</details>
---


<details><summary style="margin:3px"> Save RuleSet</summary>

<p style="margin:10px">
    <span style="background-color:rgb(0, 122, 202); color:white; padding:2px"> 
        POST
    </span>
    <code> /v1/conditions </code>
</p>

---



### ***Request Body Schema***: 
```json
{
  "id": 123,
  "name": "Kale"
}
```



### Request Fields: 
| Field | Required | Description |
| ------ | ------ | ------ |
|id|true|the user id|
|name|false|the user name|


### Response
```json
{"status": 201, "message": "created"}
```



</details>
---

