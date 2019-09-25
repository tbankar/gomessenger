import React from 'react';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import MenuItem from '@material-ui/core/MenuItem';
import TextField from '@material-ui/core/TextField';


const useStyles = makeStyles(theme => ({
    root: {
      flexGrow: 1,
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
  }));


export default function OutlinedTextFields() {
    const classes = useStyles();
    const [values, setValues] = React.useState({
      name: 'Cat in the Hat',
      age: '',
      multiline: 'Controlled',
      currency: 'EUR',
    });
  
    const handleChange = name => event => {
      setValues({ ...values, [name]: event.target.value });
    };
    return (
        <form className={classes.container} noValidate autoComplete="off">
            <TextField
            id="outlined-multiline-static"
            label=""
            multiline
            rows="40"
            defaultValue=""
            className={classes.textField}
            margin="normal"
            InputProps = {{
                readOnly:true,
            }}
            variant="outlined"
            />
            <TextField
            id="outlined-email-input"
            label="tb2"
            className={classes.textField}
            type="user"
            name="user"
            autoComplete="user"
            margin="normal"
            variant="outlined"
          />
        </form>
    );
}