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
INSERT INTO cakes (
	title, description, rating, image
)
VALUES (
	:title, :description, :rating, :image
);
`

const CAKES_REPOSITORY_UPDATE_SQL = `
UPDATE cakes 
SET title = :title, 
	description = :description,
	rating = :rating, 
	image = :image, 
    updated_at = NOW() 
WHERE id = :id AND deleted_at IS NULL
`

const CAKES_REPOSITORY_DELETE_SQL = `
UPDATE cakes 
SET deleted_at = NOW()
WHERE id = ?
`
