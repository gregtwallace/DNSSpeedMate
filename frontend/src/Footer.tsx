import { type FC } from 'react';

import { Box, Typography } from '@mui/material';
import { grey } from '@mui/material/colors';

// no props

// component
const Footer: FC = () => {
  return (
    <Box
      component='footer'
      sx={{
        py: 1,
        zIndex: (theme) => theme.zIndex.drawer + 1,
        backgroundColor: grey[300],
      }}
    >
      <Typography variant='body2' align='center'>
        &copy; 2024 Greg T. Wallace
      </Typography>
    </Box>
  );
};

export default Footer;
