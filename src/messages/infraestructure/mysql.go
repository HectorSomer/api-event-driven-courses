package infraestructure

import (
	"api_event_driven_2/src/config"
	"api_event_driven_2/src/messages/domain/entities"
	"fmt"
	"log"
)

type MySql struct {
	conn *config.Conn_MySQL
}

func NewMySql() *MySql{
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySql{conn: conn}
}

func (mysql *MySql) CreateMessage(message entities.Message) (*entities.Message, error) {
	query := "INSERT INTO message (personEmit, message, idUserTeacher) VALUES (?, ?,?)"
    result, err := mysql.conn.ExecutePreparedQuery(query, message.PersonEmit, message.Message, message.IDUserTeacher)
    if err != nil {
		fmt.Println("Error in insert course: ", err)
        return nil, err
    }
    if result != nil {
		rowsAffected,_ := result.RowsAffected()
		if rowsAffected == 1 {
			lastInserdId, err := result.LastInsertId()
			if err != nil {
				fmt.Println("Error in insert course: ", err)
				return nil, err
			}
			message.ID = int(lastInserdId)
            return &message, nil
        }
	}else{
		log.Printf("[MySQL] - Ha habido un error en la consulta (ningún resultado).")
	}
    return &message, nil
}