
export namespace time {

	/**
	 * format time
	 * @param time 
	 * @param format 
	 * @returns string
	 */
	export function formatTime(time: string | number | Date, format: string) {
		const date = new Date(time);
		const year = '0' + date.getFullYear();
		const month = '0' + (date.getMonth() + 1);
		const day = '0' + date.getDate();
		const hour = '0' + date.getHours();
		const minute = '0' + date.getMinutes();
		const second = '0' + date.getSeconds();

		if (format.includes('y')) {
			const match = format.match(/y+/);
			const length = match ? match[0].length : 0;
			format = format.replace(/y+/, year.slice(-(length)));
		}
		if (format.includes('M')) {
			const match = format.match(/M+/);
			const length = match ? match[0].length : 0;
			format = format.replace(/M+/, month.slice(-(length)));
		}
		if (format.includes('d')) {
			const match = format.match(/d+/);
			const length = match ? match[0].length : 0;
			format = format.replace(/d+/, day.slice(-(length)));
		}
		if (format.includes('H')) {
			const match = format.match(/H+/);
			const length = match ? match[0].length : 0;
			format = format.replace(/H+/, hour.slice(-(length)));
		}
		if (format.includes('m')) {
			const match = format.match(/m+/);
			const length = match ? match[0].length : 0;
			format = format.replace(/m+/, minute.slice(-(length)));
		}
		if (format.includes('s')) {
			const match = format.match(/s+/);
			const length = match ? match[0].length : 0;
			format = format.replace(/s+/, second.slice(-(length)));
		}
		return format;
	}
}