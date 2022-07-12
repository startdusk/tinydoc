export interface Tree {
	id: number;
	paraentId: number;
	label: string;
	fileType: number;
	children?: Tree[];
}