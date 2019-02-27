import { MockupList, NotificationType, SlidingPanel } from 'argo-ui';
import * as classNames from 'classnames';
import * as minimatch from 'minimatch';
import * as PropTypes from 'prop-types';
import * as React from 'react';
import { RouteComponentProps } from 'react-router';
import { Observable } from 'rxjs';

import { Autocomplete, DataLoader, EmptyState, ErrorNotification, ObservableQuery, Page, Paginate, Query } from '../../../shared/components';
import { AppContext } from '../../../shared/context';
import * as models from '../../../shared/models';
import { AppsListPreferences, AppsListViewType, services } from '../../../shared/services';
import { ApplicationCreationWizardContainer, NewAppParams, WizardStepState } from '../application-creation-wizard/application-creation-wizard';
import * as AppUtils from '../utils';
import { ApplicationsFilter } from './applications-filter';
import { ApplicationsSummary } from './applications-summary';
import { ApplicationsTable } from './applications-table';
import { ApplicationTiles } from './applications-tiles';

require('./applications-list.scss');

const APP_FIELDS = ['metadata.name', 'metadata.resourceVersion', 'metadata.creationTimestamp', 'spec', 'status.sync.status', 'status.health'];
const APP_LIST_FIELDS = APP_FIELDS.map((field) => `items.${field}`);
const APP_WATCH_FIELDS = ['result.type', ...APP_FIELDS.map((field) => `result.application.${field}`)];

function loadApplications(): Observable<models.Application[]> {
    return Observable.fromPromise(services.applications.list([], { fields: APP_LIST_FIELDS })).flatMap((applications) =>
        Observable.merge(
            Observable.from([applications]),
            services.applications.watch(null, { fields: APP_WATCH_FIELDS }).map((appChange) => {
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
            }).filter((item) => item.updated).map((item) => item.applications),
        ),
    );
}

const ViewPref = ({children}: { children: (pref: AppsListPreferences & { page: number, search: string }) => React.ReactNode }) => (
    <ObservableQuery>
        {(q) => (
            <DataLoader load={() => Observable.combineLatest(
                services.viewPreferences.getPreferences().map((item) => item.appList), q).map((items) => {
                    const params = items[1];
                    const viewPref: AppsListPreferences = {...items[0]};
                    if (params.get('proj') != null) {
                        viewPref.projectsFilter = params.get('proj').split(',').filter((item) => !!item);
                    }
                    if (params.get('sync') != null) {
                        viewPref.syncFilter = params.get('sync').split(',').filter((item) => !!item);
                    }
                    if (params.get('health') != null) {
                        viewPref.healthFilter = params.get('health').split(',').filter((item) => !!item);
                    }
                    if (params.get('namespace') != null) {
                        viewPref.namespacesFilter = params.get('namespace').split(',').filter((item) => !!item);
                    }
                    if (params.get('cluster') != null) {
                        viewPref.clustersFilter = params.get('cluster').split(',').filter((item) => !!item);
                    }
                    if (params.get('view') != null) {
                        viewPref.view = params.get('view') as AppsListViewType;
                    }
                    return {...viewPref, page: parseInt(params.get('page') || '0', 10), search: params.get('search') || ''};
            })}>
            {(pref) => children(pref)}
            </DataLoader>
        )}
    </ObservableQuery>
);

function filterApps(applications: models.Application[], pref: AppsListPreferences, search: string) {
    return applications.filter((app) =>
        (search === '' || app.metadata.name.includes(search)) &&
        (pref.projectsFilter.length === 0 || pref.projectsFilter.includes(app.spec.project)) &&
        (pref.reposFilter.length === 0 || pref.reposFilter.includes(app.spec.source.repoURL)) &&
        (pref.syncFilter.length === 0 || pref.syncFilter.includes(app.status.sync.status)) &&
        (pref.healthFilter.length === 0 || pref.healthFilter.includes(app.status.health.status)) &&
        (pref.namespacesFilter.length === 0 || pref.namespacesFilter.some((ns) => minimatch(app.spec.destination.namespace, ns))) &&
        (pref.clustersFilter.length === 0 || pref.clustersFilter.some((server) => minimatch(app.spec.destination.server, server))),
    );
}

export class ApplicationsList extends React.Component<RouteComponentProps<{}>, { appWizardStepState: WizardStepState; }> {

    public static contextTypes = {
        router: PropTypes.object,
        apis: PropTypes.object,
    };

    private get showNewAppPanel() {
        return new URLSearchParams(this.props.location.search).get('new') === 'true';
    }

    constructor(props: RouteComponentProps<{}>) {
        super(props);
        this.state = { appWizardStepState: { next: null, prev: null, nextTitle: 'Next'} };
    }

    public render() {
        return (
        <Page title='Applications' toolbar={services.viewPreferences.getPreferences().map((pref) => ({
            breadcrumbs: [{ title: 'Applications', path: '/applications' }],
            tools: (
                <React.Fragment key='app-list-tools'>
                    <div className='applications-list__view-type'>
                        <i className={classNames('fa fa-th', {selected: pref.appList.view === 'tiles'})} onClick={() => {
                            this.appContext.apis.navigation.goto('.', { view: 'tiles'});
                            services.viewPreferences.updatePreferences({ appList: {...pref.appList, view: 'tiles'} });
                        }} />
                        <i className={classNames('fa fa-th-list', {selected: pref.appList.view === 'list'})} onClick={() => {
                            this.appContext.apis.navigation.goto('.', { view: 'list'});
                            services.viewPreferences.updatePreferences({ appList: {...pref.appList, view: 'list'} });
                        }} />
                        <i className={classNames('fa fa-pie-chart', {selected: pref.appList.view === 'summary'})} onClick={() => {
                            this.appContext.apis.navigation.goto('.', { view: 'summary'});
                            services.viewPreferences.updatePreferences({ appList: {...pref.appList, view: 'summary'} });
                        }} />
                    </div>
                </React.Fragment>
            ),
            actionMenu: {
                className: 'fa fa-plus',
                items: [{
                    title: 'New Application',
                    action: () => this.setNewAppPanelVisible(true),
                }],
            },
        }))}>
            <div className='applications-list'>
                <DataLoader
                    load={() => loadApplications()}
                    loadingRenderer={() => (<div className='argo-container'><MockupList height={100} marginTop={30}/></div>)}>
                    {(applications: models.Application[]) => (
                        applications.length === 0 ? (
                            <EmptyState icon='application'>
                                <h4>No applications yet</h4>
                                <h5>Create new application to start managing resources in your cluster</h5>
                                <button className='argo-button argo-button--base' onClick={() => this.setNewAppPanelVisible(true)}>Create application</button>
                            </EmptyState>
                        ) : (
                            <div className='row'>
                                <div className='columns small-12 xxlarge-2'>
                                    <Query>
                                    {(q) => (
                                        <div className='applications-list__search'>
                                            <i className='fa fa-search'/>
                                            {q.get('search') && (
                                                <i className='fa fa-times' onClick={() => this.appContext.apis.navigation.goto('.', { search: null }, { replace: true })}/>
                                            )}
                                            <Autocomplete
                                                renderInput={(props) => (
                                                    <input {...props} onFocus={(e) => {
                                                        e.target.select();
                                                        if (props.onFocus) {
                                                            props.onFocus(e);
                                                        }
                                                    }} className='argo-field' />
                                                )}
                                                renderItem={(item) => (
                                                    <React.Fragment>
                                                        <i className='icon argo-icon-application'/> {item.label}
                                                    </React.Fragment>
                                                )}
                                                onSelect={(val) => {
                                                    this.appContext.apis.navigation.goto(`./${val}`);
                                                }}
                                                onChange={(e) => this.appContext.apis.navigation.goto('.', { search: e.target.value }, { replace: true })}
                                                value={q.get('search') || ''} items={applications.map((app) => app.metadata.name)}/>
                                        </div>
                                    )}
                                    </Query>
                                    <ViewPref>
                                    {(pref) => <ApplicationsFilter applications={applications} pref={pref} onChange={(newPref) => {
                                        services.viewPreferences.updatePreferences({appList: newPref});
                                        this.appContext.apis.navigation.goto('.', {
                                            proj: newPref.projectsFilter.join(','),
                                            sync: newPref.syncFilter.join(','),
                                            health: newPref.healthFilter.join(','),
                                            namespace: newPref.namespacesFilter.join(','),
                                            cluster: newPref.clustersFilter.join(','),
                                        });
                                    }} />}
                                    </ViewPref>
                                </div>
                                <div className='columns small-12 xxlarge-10'>
                                <ViewPref>
                                {(pref) => pref.view === 'summary' && (
                                    <ApplicationsSummary applications={filterApps(applications, pref, pref.search)} />
                                ) || (
                                    <Paginate
                                        preferencesKey='applications-list'
                                        page={pref.page}
                                        emptyState={() => (
                                            <EmptyState icon='search'>
                                                <h4>No applications found</h4>
                                                <h5>Try to change filter criteria</h5>
                                            </EmptyState>
                                        )}
                                        data={filterApps(applications, pref, pref.search)} onPageChange={(page) => this.appContext.apis.navigation.goto('.', { page })} >
                                    {(data) => (
                                        pref.view === 'tiles' && (
                                            <ApplicationTiles
                                                applications={data}
                                                syncApplication={(appName, revision) => this.syncApplication(appName, revision)}
                                                deleteApplication={(appName) => this.deleteApplication(appName)}
                                            />
                                        ) || (
                                            <ApplicationsTable applications={data}
                                                syncApplication={(appName, revision) => this.syncApplication(appName, revision)}
                                                deleteApplication={(appName) => this.deleteApplication(appName)}
                                            />
                                        )
                                    )}
                                    </Paginate>
                                )}
                                </ViewPref>
                                </div>
                            </div>
                        )
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
            if (params.directoryRecurse !== undefined) {
                source.directory = {
                    recurse: params.directoryRecurse,
                } as models.ApplicationSourceDirectory;
            }
            const syncPolicy = {} as models.SyncPolicy;
            if (params.syncPolicy !== 'manual') {
                syncPolicy.automated = {
                    prune: params.syncPolicy === 'auto-prune',
                };
            }
            await services.applications.create(params.applicationName, params.project, syncPolicy, source, {
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
