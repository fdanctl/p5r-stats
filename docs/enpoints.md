## Enpoints

- [Root](#Root)
- [User Data](#User_Data)
- [Activities](#Activities)
- [Design System](#Design_System)

### Root

#### GET `/`

Return home page. If there is user data shows profile otherwise asks to input name to create user data.

### User Data

#### POST `/user-data`

Create user data. Redirects to `/`

#### PATCH `/user-data`

Modify user data

#### DELETE `/user-data`

Deletes user data

### Activities

#### POST `/activities`

Create activity. Redirects to `/`

#### POST `/activities/[id]`

Modify activity

#### DELETE `/activities/[id]`

Deletes activity

### Design System

#### GET `/design-system`

Show test page
