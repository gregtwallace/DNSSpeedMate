import { type initDNSServersEventType } from './types/events';
import { parseInitDNSServersEventType } from './types/events';

import { useState } from 'react';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';

import { type FC } from 'react';

import { Box } from '@mui/system';

import { EventsOn } from '../wailsjs/runtime/runtime';

import InitTab from './components/Tabs/InitTab/InitTab';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function CustomTabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role='tabpanel'
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{ p: 3 }}>
          {children}
        </Box>
      )}
    </div>
  );
}

function a11yProps(index: number) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

// component
const Main: FC = () => {
  const [initResult, setInitResult] = useState<initDNSServersEventType | null>(
    null
  );

  EventsOn('init_dns_servers', (result) => {
    setInitResult(parseInitDNSServersEventType(JSON.parse(result)));
  });

  const [value, setValue] = useState(0);

  const handleChange = (_event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

  return (
    <Box
      component='main'
      sx={{
        minHeight: 0,
        flexGrow: 1,
        display: 'flex',
        flexDirection: 'column',

        // overflow: 'hidden',
        // overflowWrap: 'anywhere',
      }}
    >
      <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
        <Tabs
          value={value}
          onChange={handleChange}
          aria-label='basic tabs example'
        >
          <Tab label='Item One' {...a11yProps(0)} />
          <Tab label='Item Two' {...a11yProps(1)} />
          <Tab label='Item Three' {...a11yProps(2)} />
        </Tabs>
      </Box>
      <CustomTabPanel value={value} index={0}>
        <InitTab initResult={initResult} />
      </CustomTabPanel>
      <CustomTabPanel value={value} index={1}>
        Item Two
      </CustomTabPanel>
      <CustomTabPanel value={value} index={2}>
        Item Three
      </CustomTabPanel>
    </Box>
  );
};

export default Main;
