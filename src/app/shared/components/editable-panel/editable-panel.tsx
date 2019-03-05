import { ErrorNotification, NotificationType } from 'argo-ui';
import * as classNames from 'classnames';
import * as React from 'react';
import { Form, FormApi } from 'react-form';

import { Consumer } from '../../context';

export interface EditablePanelItem {
    title: string;
    key?: string;
    before?: React.ReactNode;
    view: string | React.ReactNode;
    edit?: (formApi: FormApi) => React.ReactNode;
}

export interface EditablePanelProps<T> {
    title?: string;
    values: T;
    validate?: (values: T) => any;
    save?: (input: T) => Promise<any>;
    items: EditablePanelItem[];
}

require('./editable-panel.scss');

export class EditablePanel<T = {}> extends React.Component<EditablePanelProps<T>, { edit: boolean; saving: boolean }> {

    private formApi: FormApi;

    constructor(props: EditablePanelProps<T>) {
        super(props);
        this.state = { edit: false, saving: false };
    }

    public render() {
        return (
            <Consumer>{(ctx) => (
            <div className={classNames('white-box editable-panel', {'editable-panel--disabled': this.state.saving})}>
                <div className='white-box__details'>
                    <div className='editable-panel__buttons'>
                        {!this.state.edit && (
                            <button onClick={() => this.setState({ edit: true })} className='argo-button argo-button--base'>Edit</button>
                        )}
                        {this.state.edit && (
                            <React.Fragment>
                                <button disabled={this.state.saving} onClick={() => !this.state.saving && this.formApi.submitForm(null)} className='argo-button argo-button--base'>
                                    Save
                                </button> <button onClick={() => this.setState({ edit: false })} className='argo-button argo-button--base-o'>Cancel</button>
                            </React.Fragment>
                        )}
                    </div>
                    {this.props.title && <p>{this.props.title}</p>}
                    {!this.state.edit && (
                        <React.Fragment>
                        {this.props.items.map((item) => (
                            <React.Fragment key={item.key || item.title}>
                            {item.before && item.before}
                            <div className='row white-box__details-row'>
                                <div className='columns small-3'>
                                    {item.title}
                                </div>
                                <div className='columns small-9'>{item.view}</div>
                            </div>
                            </React.Fragment>
                        ))}
                        </React.Fragment>
                    ) || (
                        <Form getApi={(api) => this.formApi = api} onSubmit={async (input) => {
                            try {
                                this.setState({ saving: true });
                                await this.props.save(input as any);
                                this.setState({ edit: false, saving: false });
                            } catch (e) {
                                ctx.notifications.show({
                                    content: <ErrorNotification title='Unable to save changes' e={e}/>,
                                    type: NotificationType.Error,
                                });
                            } finally {
                                this.setState({ saving: false });
                            }
                        }} defaultValues={this.props.values} validateError={this.props.validate}>
                        {(api) => (
                            <React.Fragment>
                                {this.props.items.map((item) => (
                                    <React.Fragment key={item.key || item.title}>
                                        {item.before && item.before}
                                        <div className='row white-box__details-row'>
                                            <div className='columns small-3'>
                                                {item.title}
                                            </div>
                                            <div className='columns small-9'>
                                                {item.edit && item.edit(api) || item.view}
                                            </div>
                                        </div>
                                    </React.Fragment>
                                ))}
                            </React.Fragment>
                        )}
                        </Form>
                    )}
                </div>
            </div>
            )}</Consumer>
        );
    }
}
