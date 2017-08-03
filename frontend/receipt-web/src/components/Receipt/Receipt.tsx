import * as React from 'react';

interface Props{

}

interface Result{
    fullText:string
    lang:string
    title:string
    totalPrice:string
}

interface State {
    file: File|null
    imagePreviewUrl: string
    result:Result
}

class Receipt extends React.Component<Props, State> {
    constructor(props) {
        super(props);
        this.state = {
            file: null,
            imagePreviewUrl: '',
            result: {
                fullText: '',
                lang: '',
                title: '',
                totalPrice: ''
            }
        };
    }

    _handleSubmit(e) {
        e.preventDefault();

        if (!this.state.imagePreviewUrl) {
            return
        }

        if (this.state.file) {
            console.log("_handleSubmit " + this.state.file.name)
        }

        if (!window.fetch) {
            alert("请使用更好的浏览器")
            return
        }

        let component = this

        fetch("http://127.0.0.1:8080/api/receipt/v1/receipts_extract", {
            method: "post",
            body: JSON.stringify({
                "image": {
                    "contentBase64": this.state.imagePreviewUrl.substring(this.state.imagePreviewUrl.indexOf(',') + 1)
                }
            }),
            headers: {
                'Accept': 'application/json',
                'content-type': 'application/json'
            }
        }).then(function (response) {
            return response.json()
        }).then(function (json) {
            console.log(json)
            component.setState({
                result: {
                    fullText: json.receiptInfo.fullText,
                    lang: json.receiptInfo.lang,
                    title: json.receiptInfo.title,
                    totalPrice: json.receiptInfo.totalPrice
                }
            })
            console.log(component.state)
        }).catch(function (ex) {
            console.log(ex)
        })
    }

    _handleImageChange(e) {
        e.preventDefault();

        let file = e.target.files[0];
        if (!file) {
            return
        }

        console.log("_handleImageChange " + file.name)

        let reader = new FileReader();
        reader.onloadend = () => {
            this.setState({
                file: file,
                imagePreviewUrl: reader.result
            });
            console.log(reader.result)
        }
        reader.onerror = function (error) {
            alert("错误 " + error)
        };

        reader.readAsDataURL(file)
    }

    render() {
        let {imagePreviewUrl} = this.state;
        let $imagePreview: any;
        if (imagePreviewUrl) {
            $imagePreview = (<img src={imagePreviewUrl}/>);
        }

        console.log(this.state.result.fullText)

        var fullTextStyle = {
            width: 720,
            height: 1920,
            fontSize: 20
        }

        return (
            <div>
                <form onSubmit={(e) => this._handleSubmit(e)}>
                    <input
                        type="file"
                        onChange={(e) => this._handleImageChange(e)}/>
                    <button
                        type="submit"
                        onClick={(e) => this._handleSubmit(e)}>上传图片
                    </button>
                </form>
                <div>
                    <div>
                        <p><label>语言：</label><label>{this.state.result.lang}</label></p>
                        <p><label>标题：</label><label>{this.state.result.title}</label></p>
                        <p><label>金额：</label><label>{this.state.result.totalPrice}</label></p>
                    </div>
                    <div>
                        {$imagePreview}
                    </div>
                    <div>
                        <p><textarea value={this.state.result.fullText} style={fullTextStyle}>
                        </textarea></p>
                    </div>
                </div>
            </div>
        )
    }
}

export default Receipt;
