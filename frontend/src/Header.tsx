import { type FC } from 'react';

import { AppBar, Toolbar, Typography } from '@mui/material';
import { Help } from '@mui/icons-material';

import IconButtonBrowserLink from './components/UI/Button/IconButtonBrowserLink';

// no props

// component
const Header: FC = () => {
  return (
    <AppBar
      sx={{ position: 'relative', zIndex: (theme) => theme.zIndex.drawer + 1 }}
    >
      <Toolbar variant='dense'>
        <Typography
          component='h1'
          variant='h6'
          color='inherit'
          sx={{ flexGrow: 1 }}
        >
          DNS Speed Mate
        </Typography>

        <IconButtonBrowserLink
          href='https://www.dnsspeedmate.com'
          tooltip='Help Website'
        >
          <Help />
        </IconButtonBrowserLink>
      </Toolbar>
    </AppBar>
  );
};

export default Header;
