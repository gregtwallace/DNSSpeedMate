import CssBaseline from '@mui/material/CssBaseline';
import Box from '@mui/material/Box';

import Header from './Header';
import Main from './Main';
import Footer from './Footer';

function App() {
  return (
    <div id='App'>
      <CssBaseline />

      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          height: '100vh',
          backgroundColor: (theme) =>
            theme.palette.mode === 'light'
              ? theme.palette.grey[100]
              : theme.palette.grey[900],
        }}
      >
        <Header />
        <Main />
        <Footer />
      </Box>
    </div>
  );
}

export default App;
