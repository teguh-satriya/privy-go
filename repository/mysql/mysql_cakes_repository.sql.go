package repositories

const CAKES_REPOSITORY_LIST_SQL = `
SELECT 
 cakes.id,
 cakes.title,
 cakes.description,
 cakes.rating,
 cakes.image,
 cakes.created_at,
 cakes.updated_at
FROM cakes 
WHERE cakes.deleted_at IS NULL
ORDER BY cakes.rating DESC, cakes.title;
`

const CAKES_REPOSITORY_GET_SQL = `
SELECT 
 cakes.id,
 cakes.title,
 cakes.description,
 cakes.rating,
 cakes.image,
 cakes.created_at,
 cakes.updated_at
FROM cakes 
WHERE cakes.deleted_at IS NULL AND cakes.id = ?
LIMIT 1;
`

const CAKES_REPOSITORY_CREATE_SQL = `

`

const CAKES_REPOSITORY_UPDATE_SQL = `

`

const CAKES_REPOSITORY_DELETE_SQL = `

`
