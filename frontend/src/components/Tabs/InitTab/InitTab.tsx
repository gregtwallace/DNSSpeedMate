import { type FC } from 'react';
import { type headerType } from '../../UI/TableMui/TableHeaderRow';
import { type initDNSServersEventType } from '../../../types/events';

import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';

import BoolIcon from '../../UI/BoolIcon/BoolIcon';
import TableContainer from '../../UI/TableMui/TableContainer';
import TableHeaderRow from '../../UI/TableMui/TableHeaderRow';

// table headers
const tableHeaders: headerType[] = [
  {
    id: 'ip',
    label: 'IP',
    sortable: true,
  },
  {
    id: 'hostname',
    label: 'Hostname',
    sortable: true,
  },
  {
    id: 'protocol',
    label: 'Protocol',
    sortable: false,
  },
  {
    id: 'responds_a',
    label: 'A',
    sortable: true,
  },
  {
    id: 'responds_aaaa',
    label: 'AAAA',
    sortable: true,
  },
];

type propTypes = {
  initResult: initDNSServersEventType | null;
};

const InitTab: FC<propTypes> = (props) => {
  const { initResult } = props;

  return (
    <TableContainer>
      <Table size='small'>
        <TableHead>
          <TableHeaderRow headers={tableHeaders} />
        </TableHead>
        <TableBody>
          {initResult?.dns_servers.map((serv) => (
            <TableRow key={serv.ip_and_port + '@' + serv.hostname}>
              <TableCell>{serv.ip_and_port}</TableCell>

              <TableCell>{serv.hostname}</TableCell>

              <TableCell>{serv.protocol}</TableCell>

              <TableCell>
                <BoolIcon boolVal={serv.responds_a} />
              </TableCell>

              <TableCell>
                <BoolIcon boolVal={serv.responds_aaaa} />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default InitTab;
