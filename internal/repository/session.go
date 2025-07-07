package repository

// func (repo *sqliteRepo) GetSessionByUserID(userID string) user.UserSession {

// 	ctx := context.Background()

// 	q := userQuery.New(repo.db)

// 	session, err := q.GetSessionByUserID(ctx, userID)

// 	if err != nil {
// 		fmt.Println(err.Error())

// 		return user.UserSession{}
// 	}

// 	return user.UserSession{
// 		ID:             session.ID,
// 		UserID:         session.UserID,
// 		Token:          session.Token,
// 		SessionData:    session.SessionData,
// 		CreatedAt:      session.CreatedAt,
// 		ExpiresAt:      session.ExpiresAt,
// 		LastAccessedAt: session.LastAccessedAt,
// 	}

// }

// func (repo *sqliteRepo) CreateUserSession(userID string) user.UserSession {
// 	ctx := context.Background()

// 	q := userQuery.New(repo.db)

// 	sessionID := uuid.New().String()

// 	token := uuid.New().String()

// 	createdAt := time.Now()
// 	expiresAt := createdAt.Add(2 * time.Hour)

// 	err := q.CreateSession(ctx, userQuery.CreateSessionParams{
// 		ID:          sessionID,
// 		UserID:      userID,
// 		Token:       token,
// 		SessionData: "[]",
// 		ExpiresAt:   expiresAt,
// 	})

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	return user.UserSession{
// 		ID:             sessionID,
// 		UserID:         userID,
// 		Token:          token,
// 		SessionData:    "[]",
// 		CreatedAt:      createdAt,
// 		ExpiresAt:      expiresAt,
// 		LastAccessedAt: createdAt,
// 	}

// }

// func (repo *sqliteRepo) UpdateUserSession(userID string, token string, expiresAt time.Time) user.UserSession {

// 	ctx := context.Background()

// 	q := userQuery.New(repo.db)

// 	session, err := q.UpdateSessionToken(ctx, userQuery.UpdateSessionTokenParams{
// 		Token:     token,
// 		ExpiresAt: expiresAt,
// 	})

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	return user.UserSession{
// 		ID:             session.ID,
// 		UserID:         session.UserID,
// 		Token:          session.Token,
// 		SessionData:    session.SessionData,
// 		CreatedAt:      session.CreatedAt,
// 		ExpiresAt:      session.ExpiresAt,
// 		LastAccessedAt: session.LastAccessedAt,
// 	}

// }
