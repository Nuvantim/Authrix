-- name: CreateOTP :one
INSERT INTO otp_token(code,email) VALUES ($1,$2) RETURNING *;

-- name: FindOTP :one
SELECT COUNT(*) FROM otp_token WHERE code = $1;

-- name: FindOtpByEmail :one
SELECT COUNT(*) FROM otp_token WHERE code = $1 AND email = $2;

-- name: DeleteOTP :exec
DELETE FROM otp_token WHERE email = $1;
