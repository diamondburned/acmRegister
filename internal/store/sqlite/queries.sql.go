// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: queries.sql

package sqlite

import (
	"context"
)

const deleteGuild = `-- name: DeleteGuild :execrows
DELETE FROM
	known_guilds
WHERE
	guild_id = ?
`

func (q *Queries) DeleteGuild(ctx context.Context, guildID int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteGuild, guildID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const guildInfo = `-- name: GuildInfo :one
SELECT
	guild_id, channel_id, role_id, init_user_id, registered_message
FROM
	known_guilds
WHERE
	guild_id = ?
LIMIT
	1
`

func (q *Queries) GuildInfo(ctx context.Context, guildID int64) (KnownGuild, error) {
	row := q.db.QueryRowContext(ctx, guildInfo, guildID)
	var i KnownGuild
	err := row.Scan(
		&i.GuildID,
		&i.ChannelID,
		&i.RoleID,
		&i.InitUserID,
		&i.RegisteredMessage,
	)
	return i, err
}

const initGuild = `-- name: InitGuild :exec
INSERT INTO
	known_guilds (
		guild_id,
		channel_id,
		init_user_id,
		role_id,
		registered_message
	)
VALUES
	(?, ?, ?, ?, ?)
`

type InitGuildParams struct {
	GuildID           int64
	ChannelID         int64
	InitUserID        int64
	RoleID            int64
	RegisteredMessage string
}

func (q *Queries) InitGuild(ctx context.Context, arg InitGuildParams) error {
	_, err := q.db.ExecContext(ctx, initGuild,
		arg.GuildID,
		arg.ChannelID,
		arg.InitUserID,
		arg.RoleID,
		arg.RegisteredMessage,
	)
	return err
}

const memberInfo = `-- name: MemberInfo :one
SELECT
	guild_id, user_id, metadata
FROM
	members
WHERE
	guild_id = ?
	AND user_id = ?
`

type MemberInfoParams struct {
	GuildID int64
	UserID  int64
}

func (q *Queries) MemberInfo(ctx context.Context, arg MemberInfoParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, memberInfo, arg.GuildID, arg.UserID)
	var i Member
	err := row.Scan(&i.GuildID, &i.UserID, &i.Metadata)
	return i, err
}

const registerMember = `-- name: RegisterMember :exec
INSERT INTO
	members (guild_id, user_id, metadata)
VALUES
	(?, ?, ?)
`

type RegisterMemberParams struct {
	GuildID  int64
	UserID   int64
	Metadata string
}

func (q *Queries) RegisterMember(ctx context.Context, arg RegisterMemberParams) error {
	_, err := q.db.ExecContext(ctx, registerMember, arg.GuildID, arg.UserID, arg.Metadata)
	return err
}

const unregisterMember = `-- name: UnregisterMember :execrows
DELETE FROM
	members
WHERE
	guild_id = ?
	AND user_id = ?
`

type UnregisterMemberParams struct {
	GuildID int64
	UserID  int64
}

func (q *Queries) UnregisterMember(ctx context.Context, arg UnregisterMemberParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, unregisterMember, arg.GuildID, arg.UserID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
