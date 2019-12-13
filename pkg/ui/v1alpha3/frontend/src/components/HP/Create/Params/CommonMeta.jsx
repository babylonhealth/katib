import React from 'react';
import { connect } from 'react-redux';
import makeStyles from '@material-ui/styles/makeStyles';
import Grid from '@material-ui/core/Grid';
import Tooltip from '@material-ui/core/Tooltip';
import HelpOutlineIcon from '@material-ui/icons/HelpOutline';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import DisplayNamespaceSelection from '../../../Templates/Common/DisplayNamespaceSelection';
import { changeMeta } from '../../../../actions/hpCreateActions';


const module = "hpCreate";

const useStyles = makeStyles({
    textField: {
        marginLeft: 4,
        marginRight: 4,
        width: '100%'
    },
    help: {
        padding: 4 / 2,
        verticalAlign: "middle",
        marginRight: 5,
    },
    parameter: {
        padding: 2,
        marginBottom: 10,
    },
})

const CommonParametersMeta = (props) => {
    const classes = useStyles();
    const { namespaces } = props;

    const onMetaChange = (param) => (event) => {
        props.changeMeta(param, event.target.value);
    }

    return (
        <div>
            {props.commonParametersMetadata.map((param, i) => {
                return (
                    <div key={i} className={classes.parameter}>
                        <Grid container alignItems={"center"}>
                            <Grid item xs={12} sm={3}>
                                <Typography variant={"subheading"}>
                                    <Tooltip title={param.description}>
                                        <HelpOutlineIcon className={classes.help} color={"primary"}/>
                                    </Tooltip>
                                    {param.name}
                                </Typography>
                            </Grid>
                            <Grid item xs={12} sm={8}>
                            {param.name !== "Namespace" ? (
                                <TextField
                                    className={classes.textField}
                                    value={param.value}
                                    onChange={onMetaChange(param.name)}
                                />
                            ) : (
                                <DisplayNamespaceSelection
                                    value={param.value}
                                    onChange={onMetaChange(param.name)}
                                    namespaces={namespaces}
                                />
                            )}
                            </Grid>
                        </Grid>
                    </div>
                )
            })}
        </div>
    )
}


const mapStateToProps = state => {
    return {
        commonParametersMetadata: state[module].commonParametersMetadata,
    }
}

export default connect(mapStateToProps, { changeMeta })(CommonParametersMeta);
