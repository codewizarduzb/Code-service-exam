-- create USERS table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    bio TEXT,
    website VARCHAR(255)
);

-- create POSTS table
CREATE TABLE IF NOT EXISTS posts (
    post_id UUID PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL,
    content TEXT,
    title TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    likes BIGINT DEFAULT 0,
    dislikes BIGINT DEFAULT 0,
    views BIGINT DEFAULT 0,
    media_url TEXT,
    refresh_token TEXT);

-- create COMMENTS table
CREATE TABLE IF NOT EXISTS comments (
    comment_id UUID PRIMARY KEY NOT NULL,
    post_id UUID NOT NULL,
    user_id UUID NOT NULL,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP);

-- mock data for users table
INSERT INTO users (id, username, email, password, first_name, last_name, bio, website)
VALUES
  ('550e8400-e29b-41d4-a716-446655440001', 'user1', 'user1@example.com', 'password1', 'John', 'Doe', 'Hello, I am user 1.', 'https://example.com/user1'),
  ('550e8400-e29b-41d4-a716-446655440002', 'user2', 'user2@example.com', 'password2', 'Jane', 'Smith', 'Nice to meet you!', 'https://example.com/user2'),
  ('550e8400-e29b-41d4-a716-446655440003', 'user3', 'user3@example.com', 'password3', 'Alice', 'Johnson', 'Lorem ipsum dolor sit amet.', 'https://example.com/user3'),
  ('550e8400-e29b-41d4-a716-446655440004', 'user4', 'user4@example.com', 'password4', 'Bob', 'Williams', 'Consectetur adipiscing elit.', 'https://example.com/user4'),
  ('550e8400-e29b-41d4-a716-446655440005', 'user5', 'user5@example.com', 'password5', 'Emily', 'Brown', 'Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.', 'https://example.com/user5'),
  ('550e8400-e29b-41d4-a716-446655440006', 'user6', 'user6@example.com', 'password6', 'Michael', 'Jones', 'Ut enim ad minim veniam.', 'https://example.com/user6'),
  ('550e8400-e29b-41d4-a716-446655440007', 'user7', 'user7@example.com', 'password7', 'Emma', 'Garcia', 'Quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 'https://example.com/user7'),
  ('550e8400-e29b-41d4-a716-446655440008', 'user8', 'user8@example.com', 'password8', 'William', 'Martinez', 'Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.', 'https://example.com/user8'),
  ('550e8400-e29b-41d4-a716-446655440009', 'user9', 'user9@example.com', 'password9', 'Olivia', 'Hernandez', 'Excepteur sint occaecat cupidatat non proident.', 'https://example.com/user9'),
  ('550e8400-e29b-41d4-a716-446655440010', 'user10', 'user10@example.com', 'password10', 'James', 'Young', 'Sunt in culpa qui officia deserunt mollit anim id est laborum.', 'https://example.com/user10');

-- mock data for posts table
INSERT INTO posts (post_id, user_id, content, title, likes, dislikes, views, media_url, refresh_token)
VALUES 
    ('f47ac10b-58cc-4372-a567-0e02b2c3d479', '647ac10b-58cc-4372-a567-0e02b2c3d479', 'Content 1', 'Title 1', 10, 5, 100, 'http://example.com/media1.jpg', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c'),
    ('1c8e57b3-4a30-4e42-bd79-64cb49e8c615', '3e3e5c91-d2f7-4e95-b7c7-84e14f81b0d1', 'Content 2', 'Title 2', 20, 3, 150, 'http://example.com/media2.jpg', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c'),
    ('5ab5fb6c-dc07-46cb-8ba9-34280a35856f', '7b88517b-4480-4895-a13a-29dcbfdecb19', 'Content 3', 'Title 3', 15, 8, 200, 'http://example.com/media3.jpg', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c'),
    ('aa5f7cb8-ae5f-4df3-907a-6e1c4389c8b1', 'b2d78d97-295a-4596-86f4-d08002f64834', 'Content 4', 'Title 4', 5, 2, 50, 'http://example.com/media4.jpg', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c'),
    ('bb0a3189-3e58-4f57-b7b1-8d6ac4fb326f', 'c902b5f2-063b-42b8-8e7e-415d5145f12f', 'Content 5', 'Title 5', 8, 12, 80, 'http://example.com/media5.jpg', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c');

-- mock data for comments table
INSERT INTO comments (comment_id, post_id, user_id, content)
VALUES
    ('5b73d609-ff2f-4e08-ae14-7383c15347bc', 'f47ac10b-58cc-4372-a567-0e02b2c3d479', 'df156fb2-89e5-4b0f-8d61-f04c028b8bb1', 'This is a comment on post 1.'),
    ('75f1b00d-8499-40e8-ba18-7e1290e8bcac', 'f47ac10b-58cc-4372-a567-0e02b2c3d479', 'ec3f5a3b-240b-4d1c-8307-90c51a3c4b7a', 'This is a comment on post 2.'),
    ('a8c221f3-82dc-4ce9-879e-7da42212e02d', '1c8e57b3-4a30-4e42-bd79-64cb49e8c615', 'f58263b3-c4f6-49f3-8e2d-b3c68a2d65e7', 'This is a comment on post 3.'),
    ('f510b869-70a6-4912-9823-8934b981dd2c', '1c8e57b3-4a30-4e42-bd79-64cb49e8c615', '2b18c9a3-4507-4a92-84d8-63610de631a8', 'This is a comment on post 4.'),
    ('e9f3a70b-7e86-4694-9d9d-9e674f33fe39', '5ab5fb6c-dc07-46cb-8ba9-34280a35856f', '900a5f7a-d07c-45ec-b2a4-6d82c0d9817a', 'This is a comment on post 5.'),
    ('6b73d609-fe2f-ae08-ac14-7a83c15347bc', '5ab5fb6c-dc07-46cb-8ba9-34280a35856f', 'df156fb2-89e5-4b0f-8d61-f04c028b8bb1', 'This is a comment on post 6.'),
    ('a5f1b10d-849b-40a8-bc18-6e1290e8bcac', 'aa5f7cb8-ae5f-4df3-907a-6e1c4389c8b1', 'ec3f5a3b-240b-4d1c-8307-90c51a3c4b7a', 'This is a comment on post 7.'),
    ('b7c221f3-81dc-4cb9-779e-7d242212e02d', 'aa5f7cb8-ae5f-4df3-907a-6e1c4389c8b1', 'f58263b3-c4f6-49f3-8e2d-b3c68a2d65e7', 'This is a comment on post 8.'),
    ('a510b869-7ca6-4a12-98a3-a934b981dd2c', 'bb0a3189-3e58-4f57-b7b1-8d6ac4fb326f', '2b18c9a3-4507-4a92-84d8-63610de631a8', 'This is a comment on post 9.'),
    ('a8f3a70b-6e86-4a94-8d9d-7e674f33fe39', 'bb0a3189-3e58-4f57-b7b1-8d6ac4fb326f', '900a5f7a-d07c-45ec-b2a4-6d82c0d9817a', 'This is a comment on post 10.');


-- DELETE FROM users;
-- DELETE FROM posts;
-- DELETE FROM comments;
-- DROP TABLE IF EXISTS users;
-- DROP TABLE IF EXISTS posts;
-- DROP TABLE IF EXISTS comments;