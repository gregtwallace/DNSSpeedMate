import { type FC } from 'react';

import { TableRow } from '@mui/material';

import TableHeaderSortable from './TableHeaderSortable';
import TableHeader from './TableHeader';

export type headerType = {
  id: string;
  label: string;
  sortable: boolean;
};

type propTypes = {
  headers: headerType[];
};

const TableHeaderRow: FC<propTypes> = (props) => {
  const { headers } = props;

  // setSortHandler changes the sort when a column is clicked
  const setSortHandler = (_headerId: string): void => {
  };

  return (
    <TableRow>
      {headers.map((header) => {
        if (header.sortable) {
          return (
            <TableHeaderSortable
              key={header.id}
              id={header.id}
              label={header.label}
              orderBy=''
              order='asc'
              onClick={setSortHandler}
            />
          );
        }
        return (
          <TableHeader key={header.id} id={header.id} label={header.label} />
        );
      })}
    </TableRow>
  );
};

export default TableHeaderRow;
