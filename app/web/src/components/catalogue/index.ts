export interface Tree {
	id: number;
	paraentId: number;
	label: string;
	// 1=file 2=folder
	fileType: 1 | 2;
	children?: Tree[];
	leaf?: boolean
}