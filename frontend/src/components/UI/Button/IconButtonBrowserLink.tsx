import { type FC, type ReactNode } from 'react';

import { IconButton as MuiIconButton, Tooltip } from '@mui/material';
import { BrowserOpenURL } from '../../../../wailsjs/runtime/runtime';

// prop types
type propTypes = {
  children: ReactNode;

  href: string;

  tooltip: string;
};

// component
const IconButtonBrowserLink: FC<propTypes> = (props) => {
  const { children, href, tooltip } = props;

  return (
    /* Note: Tooltip doesn't show anything if title is blank */
    <Tooltip title={tooltip}>
      <MuiIconButton onClick={() => BrowserOpenURL(href)} color='inherit'>
        {children}
      </MuiIconButton>
    </Tooltip>
  );
};
export default IconButtonBrowserLink;
