import * as React from 'react';

import * as models from '../../../shared/models';
import * as utils from '../utils';

require('./application-status-panel.scss');

export const ApplicationStatusPanel = ({application, showOperation, showConditions}: {application: models.Application, showOperation?: () => any, showConditions?: () => any}) => {
    const today = new Date();
    const creationDate = new Date(application.metadata.creationTimestamp);

    const daysActive = Math.round(Math.abs((today.getTime() - creationDate.getTime()) / (24 * 60 * 60 * 1000)));
    let daysSinceLastSynchronized = 0;
    const history = application.status.history || [];
    if (history.length > 0) {
        const deployDate = new Date(history[history.length - 1].deployedAt);
        daysSinceLastSynchronized = Math.round(Math.abs((today.getTime() - deployDate.getTime()) / (24 * 60 * 60 * 1000)));
    }
    const cntByCategory = (application.status.conditions || []).reduce(
        (map, next) => map.set(utils.getConditionCategory(next), (map.get(utils.getConditionCategory(next)) || 0) + 1),
        new Map<string, number>());
    let appOperationState = application.status.operationState;
    if (application.operation && application.status.operationState === undefined) {
        appOperationState = {
            phase:  models.OperationPhases.InProgress,
            startedAt: new Date().toISOString(),
            operation: {
                sync: {},
            } as models.Operation,
        } as models.OperationState;
    }
    return (
        <div className='application-status-panel row'>
            <div className='application-status-panel__item columns small-3'>
                <div className='application-status-panel__item-value'>{daysActive}</div>
                <div className='application-status-panel__item-name'>Days active</div>
            </div>
            <div className='application-status-panel__item columns small-3'>
                <div className='application-status-panel__item-value'>{daysSinceLastSynchronized}</div>
                <div className='application-status-panel__item-name'>Days since last synchronized</div>
            </div>
            {appOperationState && (
            <div className='application-status-panel__item columns small-3'>
                <div className={`application-status-panel__item-value application-status-panel__item-value--${appOperationState.phase}`}>
                    <a onClick={() => showOperation && showOperation()}>{utils.getOperationType(appOperationState)} <utils.OperationPhaseIcon
                        phase={appOperationState.phase}/></a>
                </div>
                <div className='application-status-panel__item-name'>
                    {appOperationState.phase} at {appOperationState.finishedAt || appOperationState.startedAt}
               </div>
            </div>
            )}
            {application.status.conditions && (
            <div className={`application-status-panel__item columns small-3`}>
                <div className='application-status-panel__item-value' onClick={() => showConditions && showConditions()}>
                    {cntByCategory.get('info') && <a className='info'>{cntByCategory.get('info')} Info</a>}
                    {cntByCategory.get('warning') && <a className='warning'>{cntByCategory.get('warning')} Warnings</a>}
                    {cntByCategory.get('error') && <a className='error'>{cntByCategory.get('error')} Errors</a>}
                </div>
            </div>
            )}
        </div>
    );
};
