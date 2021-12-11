function app() {
    return {
        cooters: {},
        interval: null,
        setup(palette, cootercount, cootersize) {
            const request = { canvasID: 'sonorous-canvas', cooterCount: parseInt(cootercount), cooterSize: parseInt(cootersize), width: 800, height: 600, palette: palette };
            const response = CooterSetup(JSON.stringify(request));
            const tas = CooterParser(response);
            return tas;
        },
        run() {
            const self = this;
            this.interval = setInterval(() => {
                const passed = self.cooters[0];
                const request = { canvasID: 'sonorous-canvas', cooters: [...passed.map(cooter => JSON.stringify(cooter))],  width: 800, height: 600 };
                const tas = CooterParser(CooterStep(JSON.stringify(request)));
                self.cooters = tas;
                return tas
            }, 50)
        },
        pause() {
            if (this.interval != null) {
                clearInterval(this.interval);
                this.interval = null;
            }
        },
        createPalettes() {
            const palettecfgs = [
                { Start: '#1B9AAA', End: '#F5F1E3', count: 10, type: 'RGB'},
                { Start: '#EE5622', End: '#221E22', count: 10, type: 'HSV'} 
            ].map(cfg => JSON.stringify(cfg));
            return CooterInterpolatedPalettes(...palettecfgs);
        },
        createSamplers() {
            const samplers = [
                { Start: '#FFD23F', End: '#95C623', count: 10, type: 'HSV'} 
            ].map(cfg => JSON.stringify(cfg));
            const plers =  CooterInterpolatedPalettes(...samplers);
            return plers[0];
        },
        drawSampler(color) {
            const request = { canvasID: 'sonorous-canvas', color: color, width: 800, height: 600};           
            DrawSampler(JSON.stringify(request));
        }
    }
}