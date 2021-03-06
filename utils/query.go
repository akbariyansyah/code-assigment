package utils

const (
	SELECT_ACCOUNT_BY_NUMBER = `SELECT 
    ma.ACCOUNT_NUMBER, mc.NAME, ma.BALANCE
FROM
    m_account ma
        JOIN
    m_customer mc ON ma.CUSTOMER_NUMBER = mc.CUSTOMER_NUMBER WHERE ma.ACCOUNT_NUMBER = ?`

	UPDATE_SENDER_BALANCE   = "UPDATE m_account SET BALANCE=BALANCE-? WHERE ACCOUNT_NUMBER=?"
	UPDATE_RECEIVER_BALANCE = "UPDATE m_account SET BALANCE=BALANCE+? WHERE ACCOUNT_NUMBER=?"
)
