package config

//DBHost :
const DBHost = "localhost"

//DBUser :
const DBUser = "postgres"

//DBPassword :
const DBPassword = "123456"

//DBName :
const DBName = "amcoop_payroll"

//DBPort :
const DBPort = 5432

// AuthSecret is the signature used in JSON Web Tokens
const AuthSecret = "I am a very dirty secret. Don't tell me to other people."

// PasswordHashingCost determines the intensity in saving our passwords
const PasswordHashingCost = 10

// TokenExpiration is 16 hours
const TokenExpiration = 16

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// Schema Names, Don't Change Unless You Have a Good Reason

// SchemaMaintenanceFile :
const SchemaGeneric = "generic"

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
