# API

## user

### FindByID

**route:** /api/user/find_by_id/{userID}

**Method:** GET

**path Parameters**

* {userID}

**request:** none

response:

```json
{
    "id": "04056053-5d96-4069-94c3-4b3281ef32a0",
    "username": "string",
    "displayName": "string",
    "currentMeetingId": "04056053-5d96-4069-94c3-4b3281ef32a0",
	"meetingHistory": [
        "04056053-5d96-4069-94c3-4b3281ef32a0",
        "04056053-5d96-4069-94c3-4b3281ef32a0",
        "04056053-5d96-4069-94c3-4b3281ef32a0"
    ],
	"rating": 1,
	"age": 1,
	"gender": 1
}
```

### Find

**route:** /api/users/find

**Method:** GET

**path Parameters**

* user_name=string
* display_name=string
* current_meeting_id=uuid
* age=int
* gender=int

**request:** none

**response:**

```json
{
    "04056053-5d96-4069-94c3-4b3281ef32a0": {
        "id": "04056053-5d96-4069-94c3-4b3281ef32a0",
        "username": "string",
        "displayName": "string",
        "currentMeetingId": "04056053-5d96-4069-94c3-4b3281ef32a0",
        "meetingHistory": [
            "04056053-5d96-4069-94c3-4b3281ef32a0",
            "04056053-5d96-4069-94c3-4b3281ef32a0",
            "04056053-5d96-4069-94c3-4b3281ef32a0"
        ],
        "rating": 1,
        "age": 1,
        "gender": 1
    },
    "04056053-5d96-4069-94c3-4b3281ef32a0": {
        ...
    }
} // это map[uuid]user
```

### Create

**route:** /api/user/create

**Method:** POST

**request:** 

```json
{
    "username": "string",
    "displayName": "string",
	"age": 1,
	"gender": 1
}
```

**response:**

```json
{
	"id": "04056053-5d96-4069-94c3-4b3281ef32a0"
}
```

### Update

**route:** /api/user/update/{userID}

**Method:** PUT

**path Parameters**

* {userID}

**request:** 

```json
{
    "username": "string",
    "displayName": "string",
	"age": 1,
	"gender": 1
}
```

**response:**

```json
{
	"id": "04056053-5d96-4069-94c3-4b3281ef32a0"
}
```

### Delete

**route:** /api/user/update/{userID}

**Method:** DELETE

**path Parameters**

* {userID}

**request:** none

**response:** none

## meeting

### FindByID

**route:** /api/meeting/find_by_id/{meetingID}

**Method:** GET

**path Parameters**

* {meetingID}

**request:** none

response:

```json
{
    "id": "04056053-5d96-4069-94c3-4b3281ef32a0",
    "gatheringPlaceId": "04056053-5d96-4069-94c3-4b3281ef32a0",
	"initiatorsId": "04056053-5d96-4069-94c3-4b3281ef32a0",
	"startTime": "2024-01-30T18:38:25.125Z",
	"endTime": "2024-01-30T18:38:25.125Z",
	"usersQuantity": 2,
	"state": 0
}
```

### Find

**route:** /api/meeting/find

**Method:** GET

**path Parameters**

* gathering place=uuid
* initiator=uuid

**request:** none

**response:**

```json
{
    "04056053-5d96-4069-94c3-4b3281ef32a0": {
        "id": "04056053-5d96-4069-94c3-4b3281ef32a0",
        "gatheringPlaceId": "04056053-5d96-4069-94c3-4b3281ef32a0",
        "initiatorsId": "04056053-5d96-4069-94c3-4b3281ef32a0",
        "startTime": "2024-01-30T18:38:25.125Z",
        "endTime": "2024-01-30T18:38:25.125Z",
        "usersQuantity": 2,
        "state": 0
    },
    "04056053-5d96-4069-94c3-4b3281ef32a0": {
        ...
    }
} // это map[uuid]user
```

### Create

**route:** /api/meeting/create

**Method:** POST

**request:** 

```json
{
    "gatheringPlaceId": "04056053-5d96-4069-94c3-4b3281ef32a0",
	"initiatorsId": "04056053-5d96-4069-94c3-4b3281ef32a0",
	"startTime": "2024-01-30T18:38:25.125Z",
	"endTime": "2024-01-30T18:38:25.125Z",
	"usersQuantity": 2,
	"state": 0
}
```

**response:**

```json
{
	"id": "04056053-5d96-4069-94c3-4b3281ef32a0"
}
```

### Update

**route:** /api/meeting/update/{meetingID}

**Method:** PUT

**path Parameters**

* {meetingID}

**request:** 

```json
{
    "gatheringPlaceId": "04056053-5d96-4069-94c3-4b3281ef32a0",
	"initiatorsId": "04056053-5d96-4069-94c3-4b3281ef32a0",
	"startTime": "2024-01-30T18:38:25.125Z",
	"endTime": "2024-01-30T18:38:25.125Z",
	"usersQuantity": 2,
	"state": 0
}
```

**response:**

```json
{
	"id": "04056053-5d96-4069-94c3-4b3281ef32a0"
}
```

### Delete

**route:** /api/meeting/update/{meetingID}

**Method:** DELETE

**path Parameters**

* {meetingID}

**request:** none

**response:** none

# gatheringPlace

### FindByID

**route:** /api/gatheringPlace/find_by_id/{gatheringPlaceID}

**Method:** GET

**path Parameters**

* {gatheringPlaceID}

**request:** none

response:

```json
{
    "id": "04056053-5d96-4069-94c3-4b3281ef32a0",
	"address": {
        "country":        "Russia",
        "city":           "Piter",
        "streetName ":    "kronverksky",
        "houseNumber":    "49",
        "BbuildingNumber" 5
    },
	"averagePrice": 1,
	"cuisineType": 0,
	"rating": 1,
	"phoneNumber": "88005553535"
}
```

### Find

**route:** /api/gatheringPlace/find

**Method:** GET

**path Parameters**

* initiator=uuid
* rating=int
* aвdress=(сложна, наверное надо разбить на несколько параметров адреса)
* cuisine_type=int

**request:** none

**response:**

```json
{
    "04056053-5d96-4069-94c3-4b3281ef32a0": {
        "id": "04056053-5d96-4069-94c3-4b3281ef32a0",
        "address": {
            "country":        "Russia",
            "city":           "Piter",
            "streetName ":    "kronverksky",
            "houseNumber":    "49",
            "BbuildingNumber" 5
        },
        "averagePrice": 1,
        "cuisineType": 0,
        "rating": 1,
        "phoneNumber": "88005553535"
    },
    "04056053-5d96-4069-94c3-4b3281ef32a0": {
        ...
    }
} // это map[uuid]user
```

### Create

**route:** /api/gatheringPlace/create

**Method:** POST

**request:** 

```json
{
	"address": {
        "country":        "Russia",
        "city":           "Piter",
        "streetName ":    "kronverksky",
        "houseNumber":    "49",
        "BbuildingNumber" 5
    },
	"averagePrice": 1,
	"cuisineType": 0,
	"rating": 1,
	"phoneNumber": "88005553535"
}
```

**response:**

```json
{
	"id": "04056053-5d96-4069-94c3-4b3281ef32a0"
}
```

### Update

**route:** /api/gatheringPlace/update/{gatheringPlaceID}

**Method:** PUT

**path Parameters**

* {gatheringPlaceID}

**request:** 

```json
{
	"address": {
        "country":        "Russia",
        "city":           "Piter",
        "streetName ":    "kronverksky",
        "houseNumber":    "49",
        "BbuildingNumber" 5
    },
	"averagePrice": 1,
	"cuisineType": 0,
	"rating": 1,
	"phoneNumber": "88005553535"
}
```

**response:**

```json
{
	"id": "04056053-5d96-4069-94c3-4b3281ef32a0"
}
```

### Delete

**route:** /api/gatheringPlace/update/{gatheringPlaceID}

**Method:** DELETE

**path Parameters**

* {gatheringPlaceID}

**request:** none

**response:** none