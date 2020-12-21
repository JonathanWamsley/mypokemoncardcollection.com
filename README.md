# A pokemon card collection tracker website

mypokemoncardcollection.com

### MVC architecture

This project uses the model-view-controller (MVC) architectural patten for organization and structure 

MVC is used because it is easy to navigate, maintain, and collaborate.

There are three main folders to separate code responsibility

1. Views
    - responsible for rendering the data
    - html code that is returned to the end user belongs here
    - try to avoid logic and focus on display
    - view layer
        - where shared layouts are stored between views
2. Controllers
    - responsible for handling most of the business logic for a web request
    - will interact with views and models, but does not create views or update databases
    - it passes data around to different parts of the application
3. Models
    - responsible for interacting with raw data in the database, APIs or other services
    - handles storing data objects as well as creating, reading, updating, deleting data

### Models


- Will use Model Struct using the gorm.Model, an ORM object-relational-mapping
- Will use ModleServices as a type that provies methods for querying, creating, and updating the model
    - uses the gorm.DB pointer
    - has an interface showing what methods it has access to

- User Credentials Model
    - description
        - holds user login details
    - columns
        - user id
        - email
        - hash pw
        - password (not stored)
        - remember (not stored)
        - remember hashed

- User Collection Model
    - description
        - holds what pokemon cards the user owns
    - columns
        - user id
        - card id
        - card set

- Pokemon Collection Model
    - description
        - holds all the pokemon cards
    - columns
        - card id
        - card name
        - card number
        - card set
        - image path


### Design pages

The webapp will be able to
- display pokemon cards in columns by set
- usec can change the set
- user can add or delete cards owned
- login/sign up

- /
    - get
    - home page
    - description and relevant opening information on the app

- /gallery
    - get
    - redirects to the latest pokemon set gallery/latest
    - show the pokemon cards based on set

- /gallery/set-name
    - get
    - drop down menu to change
    - shows the pokemon cards from set

- /gallery/set-name/mark-id
    - post
    - returns gallery of set name with update id marked
        - js change, no need to reload page
    - marks card id for the user as being owned

- /gallery/set-name/unmark-id
    - post
    - returns gallery of set name with update id unmarked
        - js change, no need to reload page
    - marks card id for the user as being owned

- /login
    - post
    - do the login

- /login
    - get
    - show login page

- /register
    - get
    - show register page

- /register
    - post
    - add user to db

- /change
    - change password