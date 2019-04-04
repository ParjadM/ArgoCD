import { Tooltip } from 'argo-ui';
import * as React from 'react';

import { Consumer } from '../../../shared/context';
import * as models from '../../../shared/models';

import { ApplicationIngressLink } from '../application-ingress-link';
import * as AppUtils from '../utils';

export interface ApplicationTilesProps {
    applications: models.Application[];
    syncApplication: (appName: string, revision: string) => any;
    deleteApplication: (appName: string) => any;
}

export const ApplicationTiles = ({applications, syncApplication, deleteApplication}: ApplicationTilesProps) => (
    <Consumer>
    {(ctx) => (
    <div className='argo-table-list argo-table-list--clickable row small-up-1 medium-up-2 large-up-4'>
        {applications.map((app) => (
            <div key={app.metadata.name} className='column column-block'>
                <div className={`argo-table-list__row
                    applications-list__entry applications-list__entry--comparison-${app.status.sync.status}
                    applications-list__entry--health-${app.status.health.status}`
                }>
                    <div className='row' onClick={(e) => ctx.navigation.goto(`/applications/${app.metadata.name}`, {}, { event: e })}>
                        <div className='columns small-12 applications-list__info'>
                            <div className='applications-list__external-link'>
                                <ApplicationIngressLink ingress={app.status.ingress}/>
                            </div>
                            <div className='row'>
                                <div className='columns applications-list__title'>{app.metadata.name}</div>
                            </div>
                            <div className='row'>
                                <div className='columns small-3'>Project:</div>
                                <div className='columns small-9'>{app.spec.project}</div>
                            </div>
                            <div className='row'>
                                <div className='columns small-3'>Status:</div>
                                <div className='columns small-9'>
                                    <AppUtils.HealthStatusIcon state={app.status.health}/> {app.status.health.status}
                                    &nbsp;
                                    <AppUtils.ComparisonStatusIcon status={app.status.sync.status}/> {app.status.sync.status}
                                </div>
                            </div>
                            <div className='row'>
                                <div className='columns small-3'>Repository:</div>
                                <div className='columns small-9'>
                                    <Tooltip content={app.spec.source.repoURL}>
                                        <a href={app.spec.source.repoURL} target='_blank' onClick={(event) => event.stopPropagation()}>
                                            <i className='fa fa-external-link'/> {app.spec.source.repoURL}
                                        </a>
                                    </Tooltip>
                                </div>
                            </div>
                            <div className='row'>
                                <div className='columns small-3'>Path:</div>
                                <div className='columns small-9'>{app.spec.source.path}</div>
                            </div>
                            <div className='row'>
                                <div className='columns small-3'>Destination:</div>
                                <div className='columns small-9'>
                                    <Tooltip content={app.spec.destination.server + '/' + app.spec.destination.namespace}>
                                        <span>{app.spec.destination.server}/{app.spec.destination.namespace}</span>
                                    </Tooltip>
                                </div>
                            </div>
                            <div className='row'>
                                <div className='columns applications-list__entry--actions'>
                                    <a className='argo-button argo-button--base'
                                        onClick={() => ctx.navigation.goto(`/applications/${app.metadata.name}`, {deploy: 'all'})}><i className='argo-icon-deploy'/> Sync</a>
                                    &nbsp;
                                    <a className='argo-button argo-button--base' onClick={() => deleteApplication(app.metadata.name)}><i className='fa fa-times-circle'/> Delete</a>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        ))}
    </div>
    )}
    </Consumer>
);
