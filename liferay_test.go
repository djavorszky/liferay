package liferay

import (
	"fmt"
	"testing"
)

func TestOracleJDBC(t *testing.T) {
	control := JDBC{
		Driver:   "jdbc.default.driverClassName=oracle.jdbc.driver.OracleDriver",
		URL:      "jdbc.default.url=jdbc:oracle:thin:@192.168.211.88:1521:orcl",
		User:     "jdbc.default.username=system",
		Password: "jdbc.default.password=password",
	}

	jdbc := OracleJDBC("192.168.211.88:1521", "orcl", "system", "password")

	if jdbc.Driver != control.Driver {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.Driver, jdbc.Driver))
	}

	if jdbc.URL != control.URL {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.URL, jdbc.URL))
	}

	if jdbc.User != control.User {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.User, jdbc.User))
	}

	if jdbc.Password != control.Password {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.Password, jdbc.Password))
	}
}

func TestPostgresJDBC(t *testing.T) {
	control := JDBC{
		Driver:   "jdbc.default.driverClassName=org.postgresql.Driver",
		URL:      "jdbc.default.url=jdbc:postgresql://192.168.211.88:5432/lportal",
		User:     "jdbc.default.username=system",
		Password: "jdbc.default.password=password",
	}

	jdbc := PostgreJDBC("192.168.211.88:5432", "lportal", "system", "password")

	if jdbc.Driver != control.Driver {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.Driver, jdbc.Driver))
	}

	if jdbc.URL != control.URL {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.URL, jdbc.URL))
	}

	if jdbc.User != control.User {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.User, jdbc.User))
	}

	if jdbc.Password != control.Password {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.Password, jdbc.Password))
	}
}

func TestMysqlJDBC(t *testing.T) {
	control := JDBC{
		Driver:   "jdbc.default.driverClassName=com.mysql.jdbc.Driver",
		URL:      "jdbc.default.url=jdbc:mysql://192.168.211.88:3306/lportal?useUnicode=true&characterEncoding=UTF-8&useFastDateParsing=false&useSSL=false",
		User:     "jdbc.default.username=root",
		Password: "jdbc.default.password=password",
	}

	jdbc := MysqlJDBC("192.168.211.88:3306", "lportal", "root", "password")

	if jdbc.Driver != control.Driver {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.Driver, jdbc.Driver))
	}

	if jdbc.URL != control.URL {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.URL, jdbc.URL))
	}

	if jdbc.User != control.User {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.User, jdbc.User))
	}

	if jdbc.Password != control.Password {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.Password, jdbc.Password))
	}
}

func TestMysqlJDBCDXP(t *testing.T) {
	control := JDBC{
		Driver:   "jdbc.default.driverClassName=com.mysql.jdbc.Driver",
		URL:      "jdbc.default.url=jdbc:mysql://192.168.211.88:3306/lportal?characterEncoding=UTF-8&dontTrackOpenResources=true&holdResultsOpenOverStatementClose=true&useFastDateParsing=false&useUnicode=true&useSSL=false",
		User:     "jdbc.default.username=root",
		Password: "jdbc.default.password=password",
	}

	jdbc := MysqlJDBCDXP("192.168.211.88:3306", "lportal", "root", "password")

	if jdbc.Driver != control.Driver {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.Driver, jdbc.Driver))
	}

	if jdbc.URL != control.URL {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.URL, jdbc.URL))
	}

	if jdbc.User != control.User {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.User, jdbc.User))
	}

	if jdbc.Password != control.Password {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.Password, jdbc.Password))
	}
}

func TestMSSQLJDBC(t *testing.T) {
	control := JDBC{
		Driver:   "jdbc.default.driverClassName=net.sourceforge.jtds.jdbc.Driver",
		URL:      "jdbc.default.url=jdbc:jtds:sqlserver//192.168.211.88:1433/master",
		User:     "jdbc.default.username=clouddb",
		Password: "jdbc.default.password=password",
	}

	jdbc := MSSQLJDBC("192.168.211.88:1433", "master", "clouddb", "password")

	if jdbc.Driver != control.Driver {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.Driver, jdbc.Driver))
	}

	if jdbc.URL != control.URL {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.URL, jdbc.URL))
	}

	if jdbc.User != control.User {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.User, jdbc.User))
	}

	if jdbc.Password != control.Password {
		t.Errorf(fmt.Sprintf("Mismatch: Expected %q, got %q", control.Password, jdbc.Password))
	}
}
