import { type FC } from 'react';

import HelpIcon from '@mui/icons-material/Help';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import DoNotDisturbOnIcon from '@mui/icons-material/DoNotDisturbOn';

type propTypes = {
  boolVal: boolean | null;
};

const BoolIcon: FC<propTypes> = (props) => {
  const { boolVal } = props;

  return (
    <>
      {boolVal == null ? (
        <HelpIcon color='info' />
      ) : boolVal ? (
        <CheckCircleIcon color='success' />
      ) : (
        <DoNotDisturbOnIcon color='error' />
      )}
    </>
  );
};

export default BoolIcon;
