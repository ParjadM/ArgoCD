import * as React from 'react';

import * as models from '../../../shared/models';

export const ApplicationConditions = ({conditions}: { conditions: models.ApplicationCondition[]}) => {
    return (
        <div>
            <h4>Application conditions</h4>
            {conditions.length == 0 && (
                <p>No conditions to report!</p>
            )}
            {conditions.length > 0 && (
            <div className='white-box'>
                <div className='white-box__details'>
                {conditions.map((condition, index) => (
                    <div className='row white-box__details-row' key={index}>
                        <div className='columns small-2'>
                            {condition.type}
                        </div>
                        <div className='columns small-10'>
                            {condition.message}
                        </div>
                    </div>
                ))}
                </div>
            </div>)}
        </div>
    );
};
