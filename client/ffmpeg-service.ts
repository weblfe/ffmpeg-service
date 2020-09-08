import * as components from "./ffmpeg-serviceComponents"

export interface ExecReq {
	action: string
	params: Array<string>
	returnBody?: string
	input: string
}

export interface ExecResp {
	msg: string
	code: number
	error?: RespError
	data: Array<components.ExecData>
	meta: components.MetaData
}

export interface RespError {
	errno: number
	errmsg: string
}

export interface CheckReq {
}

export interface CheckReqParams {
	action: string
}

export interface CheckResp {
	msg: string
	code: number
	meta?: components.MetaData
}

/**
 * @description 
 * @param req
 */
export function exec(req: ExecReq) {
	return webapi.post<ExecResp>("/exec", req)
}

/**
 * @description 
 * @param params
 */
export function check(params: CheckReqParams) {
	return webapi.get<CheckResp>("/check", params)
}
