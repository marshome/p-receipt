import * as React from 'react';
import './Receipt.css'
import CircularProgress from 'material-ui/CircularProgress';
import Textarea from 'react-textarea-autosize';

//const api_url:string='http://59.110.221.192:8081//api/receipt/v1/receipts_extract'
const api_url:string='http://127.0.0.1:8080//api/receipt/v1/receipts_extract'
const gitHubLogo=require('./GitHub-Mark-64px.png')
const googleCloudIcon=require('./GoogleCloudIcon.png')

interface Props {

}

interface Result {
    fullText: string
    lang: string
    title: string
    totalPrice: string
}

interface State {
    file: File | null
    imagePreviewUrl: string
    progressVisible: boolean
    uploadingStartTime:Date|null
    result: Result
}

class Receipt extends React.Component<Props, State> {
    timer

    constructor(props) {
        super(props);
        this.state = {
            file: null,
            imagePreviewUrl: '',
            progressVisible: false,
            uploadingStartTime: null,
            result: {
                fullText: '',
                lang: '',
                title: '',
                totalPrice: ''
            }
        };
    }

    componentDidMount() {
        this.timer = setInterval(() => {
            if (this.state.uploadingStartTime != null && !this.state.progressVisible) {
                if ((new Date().getTime() - this.state.uploadingStartTime.getTime() > 500)) {
                    this.setState({progressVisible: true})
                }
            }
        }, 50)
    }

    _handleUpload(e) {
        e.preventDefault();

        if (!this.state.imagePreviewUrl) {
            alert("请先选择图片")
            return
        }

        if (!window.fetch) {
            alert('请使用更好的浏览器');
            return
        }

        let component = this;

        component.setState({
            uploadingStartTime: new Date()
        });

        fetch(api_url, {
            method: 'post',
            body: JSON.stringify({
                'image': {
                    'contentBase64': this.state.imagePreviewUrl.substring(this.state.imagePreviewUrl.indexOf(',') + 1)
                }
            }),
            headers: {
                'Accept': 'application/json',
                'content-type': 'application/json'
            }
        }).then(function (response) {
            component.setState({progressVisible: false, uploadingStartTime: null})

            if (response.status > 299) {
                response.json().then(function (json) {
                    if (json){
                        alert(response.status+"\n"+ JSON.stringify(json,null,"  "))
                    }else{
                        alert(response.status )
                    }
                });

                component.setState({
                    result: {
                        fullText: '',
                        lang: '',
                        title: '',
                        totalPrice: ''
                    }
                });

                return
            }

            response.json().then(function (json) {
                if (!json) {
                    return
                }

                component.setState({
                    result: {
                        fullText: json.receiptInfo.fullText,
                        lang: json.receiptInfo.lang,
                        title: json.receiptInfo.title,
                        totalPrice: json.receiptInfo.totalPrice
                    }
                });
            }).catch(function (ex) {
                alert(ex)
            })
        }).catch(function (ex) {
            component.setState({progressVisible: false, uploadingStartTime: null})
            alert("网络异常: " + ex)
        })
    }

    _handleImageChange(e) {
        e.preventDefault();

        let file = e.target.files[0];
        if (!file) {
            return
        }

        if (!file.name.endsWith(".jpg") && !file.name.endsWith(".jpeg") && !file.name.endsWith(".png")) {
            alert("请选择jpg,jpeg,png格式的图片")
            return
        }

        let reader = new FileReader();
        reader.onloadend = () => {
            this.setState({
                file: file,
                imagePreviewUrl: reader.result,
                result: {
                    fullText: '',
                    lang: '',
                    title: '',
                    totalPrice: ''
                }
            });
        };
        reader.onerror = function (error) {
            alert('错误 ' + error)
        };

        reader.readAsDataURL(file)
    }

    render() {
        let fileInput: any

        let imagePreview: any;
        if (this.state.imagePreviewUrl) {
            imagePreview = (<img className="Receipt-preview" src={this.state.imagePreviewUrl}/>);
        }

        let progressView: any
        if (this.state.progressVisible) {
            progressView = (<div className="Receipt-progress">
                <CircularProgress size={60} thickness={10}/>
            </div>);
        }

        let resultView: any;
        if (this.state.result.fullText) {
            resultView = (
                <div>
                    <div>
                        <p>
                            <label className="Receipt-resultName">语言：</label>
                            <label className="Receipt-resultText">{this.state.result.lang}</label></p>
                        <p>
                            <label className="Receipt-resultName">标题：</label>
                            <label className="Receipt-resultText">{this.state.result.title}</label></p>
                        <p>
                            <label className="Receipt-resultName">金额：</label>
                            <label className="Receipt-resultText">{this.state.result.totalPrice}</label></p>
                    </div>
                    <div>
                        <Textarea className="Receipt-fullText" value={this.state.result.fullText}/>
                    </div>
                </div>);
        }

        return (
            <div className="Receipt">
                <form>
                    <a className="Receipt-fileSelector" href="javascript:void(0);" onClick={function () {
                        if (fileInput) {
                            fileInput.click()
                        }
                    }}>选择文件</a>
                    <input className="Receipt-file" type="file" ref={(input) => {
                        fileInput = input
                    }} onChange={(e) => this._handleImageChange(e)}/>
                    <button className="Receipt-upload" type="submit" onClick={(e) => this._handleUpload(e)}>上传图片
                    </button>
                    <a href="https://github.com/marshome/p-receipt"><img className="App-gitHubLogo" src={gitHubLogo}/></a>
                    <a href="https://cloud.google.com/vision/docs/"><img className="App-gitHubLogo" src={googleCloudIcon}/></a>
                </form>
                <div>{resultView}</div>
                <div>
                    {imagePreview}
                    {progressView}
                </div>
            </div>
        )
    }
}

export default Receipt;
