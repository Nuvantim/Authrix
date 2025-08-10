-- name: CreateOTP :exec
INSERT INTO otp_token(code,email) VALUES ($1,$2);

-- name: FindOTP :one
SELECT id, email FROM otp_token WHERE code = $1;

-- name: FindOtpByEmail :one
SELECT COUNT(*) FROM otp_token WHERE code = $1 AND email = $2;

-- name: DeleteOTP :exec
DELETE FROM otp_token WHERE email = $1;
