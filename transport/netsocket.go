package transport

/**
	This file is part of goPrinterRelay.

	goPrinterRelay - printer status page and protocol relay for daVinci jr 3d printers

    goPrinterRelay is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    goPrinterRelay is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with goPrinterRelay.  If not, see <http://www.gnu.org/licenses/>.

**/
import (
	"io"
	"net"
)

type netSocket struct {
	connectionStr string
	conn          net.Conn
}

func (s *netSocket) Connect() (io.Reader, io.Writer, error) {
	con, err := net.Dial("tcp", s.connectionStr)
	if err != nil {
		return nil, nil, err
	}
	s.conn = con

	return con, con, nil
}

func (s *netSocket) Disconnect() {
	s.conn.Close()
}
