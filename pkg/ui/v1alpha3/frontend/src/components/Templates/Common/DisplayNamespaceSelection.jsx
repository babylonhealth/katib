import React from 'react';
import makeStyles from '@material-ui/styles/makeStyles';
import Select from '@material-ui/core/Select';
import TextField from '@material-ui/core/TextField';
import MenuItem from '@material-ui/core/MenuItem';

const useStyles = makeStyles({
    textField: {
        marginLeft: 4,
        marginRight: 4,
        width: '100%'
    }
})

const DisplayNamespaceSelection = (props) => {
    const classes = useStyles();
    const { namespaces, onChange, value } = props;

    if (!namespaces || namespaces.length < 1) {
        return (
            <TextField
            className={classes.TextField}
            value={value}
            onChange={onChange}
            />
        )    
    } else {

    }
    return (
        <Select
            value={value !== "" ? value : namespaces[0]}
            onChange={onChange}
        >
            {namespaces.map((ns, i) => {
                return <MenuItem key={"ns-"+i} value={ns}>{ns}</MenuItem>
            })}
        </Select>
    )
}

export default DisplayNamespaceSelection
