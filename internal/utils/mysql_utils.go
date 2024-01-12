package utils

import (
	"encoding/base64"
	"encoding/json"
	"time"
    "log"

	"github.com/devanfer02/litecartes/domain"
)

func CurrentTime() string {
    return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func GetCursor(req *domain.PaginationRequest) (domain.Cursor, error) {
    var cursor domain.Cursor

    if req == nil {
        return domain.Cursor{
            LimitData: 10,
            PointNext: true,
            CreatedAt: "",
        }, nil 
    } 

    cursor, err := DecodeCursor(req.EncodedCursor)
    if err != nil {
        log.Printf("failed to decode cursor")
        return domain.Cursor{}, domain.ErrBadRequest
    }

    if req.Limit != 0 {
        cursor.LimitData = req.Limit
    }

    if !req.PointNext {
        cursor.PointNext = req.PointNext
    }

    return cursor, nil 
}

func GetPaginationOperator(pointNext bool) string {
    if pointNext {
        return ">"
    }

    return "<"
}

func CreateCursor(createdAt time.Time, next bool, limit int) domain.Cursor {
    return domain.Cursor{
        CreatedAt: createdAt.Format("2006-01-02 15:04:05"),   
        PointNext: next,
        LimitData: limit,
    }
}

func CreatePaginationResponse(prev domain.Cursor, next domain.Cursor) domain.PaginationResponse {
    return domain.PaginationResponse{
        PrevCursor: encodeCursor(prev),
        NextCursor: encodeCursor(next),
        LimitData: next.LimitData,
    }
}

func encodeCursor(cursor domain.Cursor) string {
    if cursor == (domain.Cursor{}) {
        return ""
    }

    serializedCursor, err := json.Marshal(cursor)
    if err != nil {
        return ""
    }

    encoded := base64.StdEncoding.EncodeToString(serializedCursor)

    return encoded 
}

func DecodeCursor(cursor string) (domain.Cursor, error) {
    if (cursor == "") {
        return domain.Cursor{
            CreatedAt: "",
        }, nil 
    }

    decoded, err := base64.StdEncoding.DecodeString(cursor)
    if err != nil {
        return domain.Cursor{}, err 
    }

    var crs domain.Cursor
    if err := json.Unmarshal(decoded, &crs); err != nil {
        return domain.Cursor{}, err 
    }

    return crs, nil
}