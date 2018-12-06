import { DropDownMenu, MockupList, NotificationType, SlidingPanel, TopBarFilter } from 'argo-ui';
import * as moment from 'moment';
import * as PropTypes from 'prop-types';
import * as React from 'react';
import { RouteComponentProps } from 'react-router';
import { Observable } from 'rxjs';

import { DataLoader, ErrorNotification, Page } from '../../../shared/components';
import { AppContext } from '../../../shared/context';
import * as models from '../../../shared/models';
import { services } from '../../../shared/services';
import { ApplicationCreationWizardContainer, NewAppParams, WizardStepState } from '../application-creation-wizard/application-creation-wizard';
import * as AppUtils from '../utils';

require('./applications-list.scss');

const APP_FIELDS = ['metadata.name', 'metadata.resourceVersion', 'spec', 'status.sync.status', 'status.health'];
const APP_LIST_FIELDS = APP_FIELDS.map((field) => `items.${field}`);
const APP_WATCH_FIELDS = ['result.type', ...APP_FIELDS.map((field) => `result.application.${field}`)];

type Props = RouteComponentProps<{}>;
interface State {
    allProjects: string[];
    appWizardStepState: WizardStepState;
}

export class ApplicationsList extends React.Component<Props, State> {

    public static contextTypes = {
        router: PropTypes.object,
        apis: PropTypes.object,
    };

    private get showNewAppPanel() {
        return new URLSearchParams(this.props.location.search).get('new') === 'true';
    }

    constructor(props: Props) {
        super(props);
        this.state = { allProjects: [], appWizardStepState: { next: null, prev: null, nextTitle: 'Next'} };
    }

    public async componentDidMount() {
        this.setState({ allProjects: await services.projects.list().then((items) => items.map((item) => item.metadata.name)) });
    }

    public render() {
        const projectsFilter = new URLSearchParams(this.props.history.location.search).getAll('proj');
        const filter: TopBarFilter<string> = {
            items: this.state.allProjects.map((proj) => ({ value: proj, label: proj })),
            selectedValues: projectsFilter,
            selectionChanged: (items) => {
                this.appContext.apis.navigation.goto('.', { proj: items});
            },
        };
        return (
        <Page title='Applications' toolbar={{
                filter,
                breadcrumbs: [{ title: 'Applications', path: '/applications' }],
                actionMenu: {
                    className: 'fa fa-plus',
                    items: [{
                        title: 'New Application',
                        action: () => this.setNewAppPanelVisible(true),
                    }],
                },
            }} >
                <div className='argo-container applications-list'>
                    <DataLoader
                        input={projectsFilter}
                        load={(projects) =>
                            Observable.fromPromise(services.applications.list(projects, { fields: APP_LIST_FIELDS })).flatMap((applications) =>
                            Observable.merge(
                                Observable.from([applications]),
                                services.applications.watch(null, { fields: APP_WATCH_FIELDS }).filter((appChange) =>
                                    projects.length === 0 || projectsFilter.includes(appChange.application.spec.project),
                                ).map((appChange) => {
                                    const index = applications.findIndex((item) => item.metadata.name === appChange.application.metadata.name);
                                    if (index > -1 && appChange.application.metadata.resourceVersion === applications[index].metadata.resourceVersion) {
                                        return {applications, updated: false};
                                    }
                                    switch (appChange.type) {
                                        case 'DELETED':
                                            if (index > -1) {
                                                applications.splice(index, 1);
                                            }
                                            break;
                                        default:
                                            if (index > -1) {
                                                applications[index] = appChange.application;
                                            } else {
                                                applications.unshift(appChange.application);
                                            }
                                            break;
                                    }
                                    return {applications, updated: true};
                                }).filter((item) => item.updated).map((item) => item.applications)),
                            )
                        }
                        loadingRenderer={() => <MockupList height={300} marginTop={30}/>}>
                        {(applications: models.Application[]) => (
                            <div className='argo-table-list argo-table-list--clickable row small-up-1 large-up-2'>
                                {applications.map((app) => (
                                    <div key={app.metadata.name} className='column column-block'>
                                        <div className={`argo-table-list__row
                                            applications-list__entry applications-list__entry--comparison-${app.status.sync.status}
                                            applications-list__entry--health-${app.status.health.status}`
                                        }>
                                            <div className='row' onClick={(e) => this.appContext.apis.navigation.goto(`/applications/${app.metadata.name}`, {}, e)}>
                                                <div className='columns small-12 applications-list__info'>
                                                    <div className='row'>
                                                        <div className='columns applications-list__title'>{app.metadata.name}</div>
                                                    </div>
                                                    <div className='row'>
                                                        <div className='columns small-3'>Project:</div>
                                                        <div className='columns small-9'>{app.spec.project}</div>
                                                    </div>
                                                    <div className='row'>
                                                        <div className='columns small-3'>Namespace:</div>
                                                        <div className='columns small-9'>{app.spec.destination.namespace}</div>
                                                    </div>
                                                    <div className='row'>
                                                        <div className='columns small-3'>Cluster:</div>
                                                        <div className='columns small-9'>{app.spec.destination.server}</div>
                                                    </div>
                                                    <div className='row'>
                                                        <div className='columns small-3'>Status:</div>
                                                        <div className='columns small-9'>
                                                            <AppUtils.ComparisonStatusIcon status={app.status.sync.status}/> {app.status.sync.status}
                                                        </div>
                                                    </div>
                                                    <div className='row'>
                                                        <div className='columns small-3'>Health:</div>
                                                        <div className='columns small-9'>
                                                            <AppUtils.HealthStatusIcon state={app.status.health}/> {app.status.health.status}
                                                        </div>
                                                    </div>
                                                    <div className='row'>
                                                        <div className='columns small-3'>Age:</div>
                                                        <div className='columns small-9'>{this.daysBeforeNow(app.metadata.creationTimestamp)} days</div>
                                                    </div>
                                                    <div className='row'>
                                                        <div className='columns small-3'>Repository:</div>
                                                        <div className='columns small-9'>
                                                            <a href={app.spec.source.repoURL} target='_blank' onClick={(event) => event.stopPropagation()}>
                                                                <i className='fa fa-external-link'/> {app.spec.source.repoURL}
                                                            </a>
                                                        </div>
                                                    </div>
                                                    <div className='row'>
                                                        <div className='columns small-3'>Path:</div>
                                                        <div className='columns small-9'>{app.spec.source.path}</div>
                                                    </div>
                                                    <div className='row'>
                                                        <div className='columns applications-list__entry--actions'>
                                                            <DropDownMenu anchor={() =>
                                                                <button className='argo-button argo-button--base-o'>Actions  <i className='fa fa-caret-down'/></button>
                                                            } items={[
                                                                { title: 'Sync', action: () => this.syncApplication(app.metadata.name, app.spec.source.targetRevision || 'HEAD') },
                                                                { title: 'Delete', action: () => this.deleteApplication(app.metadata.name) },
                                                            ]} />
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                ))}
                            </div>
                        )}
                    </DataLoader>
                </div>
                <SlidingPanel isShown={this.showNewAppPanel} onClose={() => this.setNewAppPanelVisible(false)} header={
                        <div>
                            <button className='argo-button argo-button--base'
                                    disabled={!this.state.appWizardStepState.prev}
                                    onClick={() => this.state.appWizardStepState.prev && this.state.appWizardStepState.prev()}>
                                Back
                            </button> <button className='argo-button argo-button--base'
                                    disabled={!this.state.appWizardStepState.next}
                                    onClick={() => this.state.appWizardStepState.next && this.state.appWizardStepState.next()}>
                                {this.state.appWizardStepState.nextTitle}
                            </button> <button onClick={() => this.setNewAppPanelVisible(false)} className='argo-button argo-button--base-o'>
                                Cancel
                            </button>
                        </div>
                    }>
                    {this.showNewAppPanel && <ApplicationCreationWizardContainer
                        createApp={(params) => this.createApplication(params)}
                        onStateChanged={(appWizardStepState) => this.setState({ appWizardStepState })} />
                    }
                </SlidingPanel>
            </Page>
        );
    }

    private get appContext(): AppContext {
        return this.context as AppContext;
    }

    private setNewAppPanelVisible(isVisible: boolean) {
        this.appContext.router.history.push(`${this.props.match.url}?new=${isVisible}`);
    }

    // DaysBeforeNow returns the delta, in days, between now and a given timestamp.
    private daysBeforeNow(timestamp: string): number {
        const end = moment();
        const start = moment(timestamp);
        const delta = moment.duration(end.diff(start));
        return Math.round(delta.asDays());
    }

    private async createApplication(params: NewAppParams) {
        try {
            const source = {
                path: params.path,
                repoURL: params.repoURL,
                targetRevision: params.revision,
                componentParameterOverrides: null,
            } as models.ApplicationSource;
            if (params.valueFiles) {
                source.helm = {
                    valueFiles: params.valueFiles,
                } as models.ApplicationSourceHelm;
            }
            if (params.environment) {
                source.ksonnet = {
                    environment: params.environment,
                } as models.ApplicationSourceKsonnet;
            }
            if (params.namePrefix) {
                source.kustomize = {
                    namePrefix: params.namePrefix,
                } as models.ApplicationSourceKustomize;
            }
            await services.applications.create(params.applicationName, params.project, source, {
                server: params.clusterURL,
                namespace: params.namespace,
            });
            this.setNewAppPanelVisible(false);
        } catch (e) {
            this.appContext.apis.notifications.show({
                content: <ErrorNotification title='Unable to create application' e={e}/>,
                type: NotificationType.Error,
            });
        }
    }

    private async syncApplication(appName: string, revision: string) {
        const synced = await AppUtils.syncApplication(appName, revision, false, false, null, this.appContext);
        if (synced) {
            this.appContext.apis.notifications.show({
                type: NotificationType.Success,
                content: `Synced revision`,
            });
        }
    }

    private async deleteApplication(appName: string) {
        const deleted = await AppUtils.deleteApplication(appName, this.appContext);
        if (deleted) {
            this.appContext.router.history.push('/applications');
        }
    }
}
